# import from the f_ku_orn_cov.json file and make a dictionary of the data


import json
import pathlib

# Path to the data file

specimen = []

with open(pathlib.Path(__file__).parent / 'orn_import.json') as f:
    data = json.load(f)
    print(data.keys())
    specimen = data['entries']


print(len(specimen))
