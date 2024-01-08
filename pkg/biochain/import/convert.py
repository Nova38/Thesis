# Load the ku_orn_database_great_plains_pre_1970_NoDups.csv file and convert it to a json file

import betterproto

import csv
import json
import uuid
from rich import print

UUID_NAMESPACE = uuid.uuid5(uuid.NAMESPACE_DNS, "ku_orn.specimen.biochain.iitc.ku.edu")

data_list = []
specimen_list = []

sex_list = set()
age_list = set()


row_mapping = {
    "primary.catalogDate.verbatim": "primary.catalogDate",
    "primary.originalDate.verbatim": "primary.originalDate",
    "georeference.georeferencedDate.verbatim": "georeference.georeferencedDate",
    "secondary.preparation.initial.verbatim": "secondary.preparation",
}

sex_cov = {
    1: ["unknown", "U", "u"],
    2: ["HERMAPHRODI"],
    3: ["M", "m", "male"],
    4: ["female", "F", "f", "F?"],
}

age_cov = {
    # AGE_UNDEFINED = 0;
    0: [],
    # AGE_UNKNOWN = 1;
    1: ["Unknown"],
    # AGE_NEST = 2;
    2: ["NESTLING", "nestling", "nestling; Immature"],
    # AGE_EMBRYO_EGG = 3;
    3: [],
    # AGE_CHICK_SUBADULT = 4;
    4: [
        "chick",
        "ca 12 hours; Downy chick",
        "11 days; Downy chick",
        "ca 80 hours; Downy chick",
        "ca 85 hours; Downy chick",
        "ca 60 hours; Downy chick",
        "Downy chick",
        "7 days; Downy chick",
        "JUST HATCHED; Downy chick",
        "ca 36 hours; Downy chick",
        "22 days; Downy chick",
        "8 days; Downy chick",
        "15 days; Downy chick",
        "LESS THAN ONE DAY OLD; Downy chick",
        "4 days; Downy chick",
        "10 days; Downy chick",
        "6 days; Downy chick",
        "chick; 2 days old",
        "66 hours; Downy chick",
        "29 days; Downy chick",
    ],
    # AGE_ADULT = 5;
    5: [
        "adult",
        "ad",
        "Ad" " adult",
        "ADULT",
        " ADULT",
        "snco; Adult",
        "Adult",
        "ADULT BY PLUMAGE",
        "ADULT, SKULL COMPLETELY OSSIFIED",
        "ADULT by plumage",
        "adult?",
        " adult",
        "subadult",
        "adult?",
        "nearst adult",
        "ADULT",
        "ADULT by plumage",
        "ADULT; SKULL COMPLETELY OSSIFIED",
        '"immature" on tag; Adult',
        "11 MONTHS OLD, BASED ON BANDING; Adult",
    ],
    # AGE_CONTINGENT = 6;
    6: [],
}

float_conversion = [
    "georeference.latitude",
    "georeference.longitude",
    "secondary.weight",
]
int_conv = ["georeference.coordinateUncercaintyInMeters"]


# Open the csv file
with open("ku_orn_database_great_plains_pre_1970_NoDups.csv", "r") as csv_file:
    data = csv.DictReader(
        csv_file,
    )

    for row in data:
        for key, value in row.items():
            if value in ["NA", "N/A"]:
                row[key] = ""

        row = {key: value for key, value in row.items() if value != ""}
        # row = {key: value for key, value in row.items() if value != }
        for keys in float_conversion:
            if keys in row:
                try:
                    row[keys] = float(row[keys])
                except:
                    print(f"Invalid Float Conversion: {keys=} - {row[keys]=}")
                    del row[keys]

        for keys in int_conv:
            if keys in row:
                try:
                    row[keys] = int(float(row[keys]))
                except:
                    print(f"Invalid Int Conversion {row[keys]=}")
                    del row[keys]

        for key, value in row_mapping.items():
            if value in row:
                row[key] = row[value]
                del row[value]

        # secondary conv
        note = ""

        raw_sex = ""
        if "secondary.sex" in row:
            raw_sex = row["secondary.sex"]

        raw_age = ""
        if "secondary.age" in row:
            raw_age = row["secondary.age"]

        if "secondary.notes" in row:
            note = row["secondary.notes"]
        if "secondary.sex" in row:
            note = f"secondary.sex converted from: {raw_sex};" + note
        if "secondary.age" in row:
            note = f"secondary.age converted from: {raw_age};" + note
        if note != "":
            row["secondary.notes"] = note

        for key, value in sex_cov.items():
            if raw_sex in value:
                row["secondary.sex"] = key
                break
        else:
            row["secondary.sex"] = 0
            sex_list.add(raw_sex)

        for key, value in age_cov.items():
            if raw_age in value:
                row["secondary.age"] = key
                break
        else:
            age_list.add(raw_age)
            row["secondary.age"] = 0

        row["collectionId"] = "ku_orn"
        row["specimenId"] = str(uuid.uuid5(UUID_NAMESPACE, row["index"]))

        data_list.append(row)


print(f"{sex_list=}")

print(f"{age_list=}")

with open("ku_orn_database_great_plains_pre_1970_NoDups.json", "w") as json_file:
    # Write the list to the json file
    json_file.write(json.dumps({"items": data_list}, indent=4))
