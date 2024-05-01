# %%

# import from the f_ku_orn_cov.json file and make a dictionary of the data
import rich
from rich import inspect
from rich.progress import track


import json
import pathlib
import requests

# Path to the data file
import more_itertools
import logging
from rich.logging import RichHandler

# logging.basicConfig(
#     level="NOTSET",
#     format="%(message)s",
#     datefmt="[%X]",
#     handlers=[RichHandler(rich_tracebacks=True)],
# )
# %%
specimen = []

with open(pathlib.Path(__file__).parent / "f_ku_orn_cov.json") as f:
    data = json.load(f)
    print(data.keys())
    specimen = data["entries"]


x = [i for i in specimen if i['specimenId'] == "f70cb713-08a7-5c7c-b58d-b5f83a3e93d7"]
re = specimen[specimen.index(x[0])+1:]


# %%


s = requests.Session()


s.headers.update(
    {
        "x-h3-session": "Fe26.2**72d10cfdb514df18749c30fbe29ffc01d47e0baf4fc982614a03fd104f82967c*hEfVeWdO6HCDv4b4CWbnYQ*2XyBVo7hsYMkAGvy3-Zzit4AbHlvs7zmvGVdVq92_nktROGPfsIW9Su9INNHP7-Ak0GWGccQhuGRF-xT_UKPmf2duwuIzIw7TmFmrJNaMbX3YBlpbvRMyugzimH0G04dd5Ni0cfPCHcvUwZr8MzXuQ**3680b31c6e5207eb13a186b9024187302074f185f054c17fc57ac638917530c1*3-6TcZy5Dhs4rgTEJUulbK4AOYRu8E1k7Ir6RVjDBws"
    }
)
r = s.get("https://biochain.ittc.ku.edu/api/auth/session")
inspect(r)

s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
    params={"collectionId": "KU Ornithology Great Plains"},
    json={"entries": re[:10]},
)


# %%
r = s.post(
    "http://localhost:8000/api/auth/login",
    json={"username": "test", "passwosrd": "tessst"},
)
print(r.cookies.items())

r = s.get("http://localhost:8000/api/auth/session")
print(r.json())

# %%

# print(s.get("http://localhost:8000/api/cc/roles/currentUser").text)

replies = []

groups = [a for a in more_itertools.batched(specimen, 100)]

for g in track(groups[1:], description="Sending batches"):
    replies.append(s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
        #params={"collectionId": ""},
        params={"collectionId": "KU Ornithology Great Plains"},
        json={"entries": g},
    )
    )

# %%
# load the specimen


print(len(specimen))

group = specimen[:200]

replies = []




for g in track(more_itertools.batched(group, 50), description="Sending batches"):
    replies.append(s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
        #params={"collectionId": ""},
        params={"collectionId": "KU Ornithology Great Plains"},
        json={"entries": g},
    )
    )


# %%
s = requests.Session()

r = s.post(
    "http://localhost:8000/api/cc/bootstrap/initSpecimens",
    params={"collectionId": "KU Ornithology Great Plains"},
    json={"entries": group},
)
print(r.text)
# %%


for g in track(batch[10:], description="Sending batches"):
    replies.append(s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
        #params={"collectionId": ""},
        params={"collectionId": "KU Ornithology Great Plains"},
        json={"entries": g},
    )
    )





s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
    params={"collectionId": "KU Ornithology Great Plains"},
    json={"entries": re[:10]},
)



s.post(
        "https://biochain.ittc.ku.edu/api/cc/bootstrap/initSpecimens",
    params={"collectionId": "KU Ornithology Great Plains"},
    json={"entries": re[:10]},
)
