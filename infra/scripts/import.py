# %%

# import from the f_ku_orn_cov.json file and make a dictionary of the data
import rich
from rich import inspect


import json
import pathlib
import requests
# Path to the data file
import more_itertools

# %%
s = requests.Session()


s.cookies.set("nuxt-session", "Fe26.2**c0f79cc0454aab1f284b6a7f5993b1793ea097da26492e6bc84c826b0f1cbfe7*549TFrgxMa_0J2n569dtpw*sWjij10CGjifXqg9dQlua3BUwK9CDcD4eRyk7Bt7_x7hosUl0sIoo6pd-aeM7q1wyaieVOAp_ctQuSOQpbLoFcmEA1NtbeOzJlmsdqgqiwyKpsuIF-livpThfNtZ6Q3oIXE_PSfOM2ovK2QubW7LRg**e446940e172f62becac7d202ac40211821d4b64e5078b1350edb87d85b273275*kSoLmWe0RIkV_Pap1DQXGEqkcVJ1sqchQQPiXpwWz6w")


s.headers.update({
    "x-h3-session": "Fe26.2**c0f79cc0454aab1f284b6a7f5993b1793ea097da26492e6bc84c826b0f1cbfe7*549TFrgxMa_0J2n569dtpw*sWjij10CGjifXqg9dQlua3BUwK9CDcD4eRyk7Bt7_x7hosUl0sIoo6pd-aeM7q1wyaieVOAp_ctQuSOQpbLoFcmEA1NtbeOzJlmsdqgqiwyKpsuIF-livpThfNtZ6Q3oIXE_PSfOM2ovK2QubW7LRg**e446940e172f62becac7d202ac40211821d4b64e5078b1350edb87d85b273275*kSoLmWe0RIkV_Pap1DQXGEqkcVJ1sqchQQPiXpwWz6w"
})
r = s.get("http://localhost:8000/api/auth/session")
inspect(r)



# %%
r = s.post(
    "http://localhost:8000/api/auth/login",
    json={"username": "test", "passwosrd": "tessst"},
)
print(r.cookies.items())

r = s.get("http://localhost:8000/api/auth/session")
print(r.json())

# %%

print(s.get("http://localhost:8000/api/cc/roles/currentUser").text)

# %%
# load the specimen
specimen = []

with open(pathlib.Path(__file__).parent / "f_ku_orn_cov.json") as f:
    data = json.load(f)
    print(data.keys())
    specimen = data["entries"]


print(len(specimen))

group = specimen[:200]

for g in more_itertools.batched(group, 200):
    print("sending new batch")
    r = s.post(
    "http://localhost:8000/api/cc/bootstrap/initSpecimens?collectionId=Chain",
        json={"entries": g},
    )
    print(r)
    print(r.text)

# %%
s = requests.Session()

r = s.post(
    "http://localhost:8000/api/cc/bootstrap/initSpecimens?collectionId=Chain",
    json={"entries": group},
)
print(r.text)
# %%
