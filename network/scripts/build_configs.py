import pathlib
import sys
import yaml
import rich
from dataclasses import dataclass
from typing import List

from rich import print


@dataclass
class User:
    name: str
    private_key: pathlib.Path
    certificate: pathlib.Path


def GetBasePath() -> pathlib.Path:
    """Returns the base path of the project."""
    return pathlib.Path(__file__).parent.parent

def GetCaliperPath() -> pathlib.Path:
    """Returns the path to the caliper network file."""
    return pathlib.Path(__file__).parent.parent.parent / "caliper-workspace/networks"/"test-network.yaml"



def MakeYaml(users: List[User]):
    """Creates a YAML file with an entry for each user in the identity."""

    data = {}
    with open(GetCaliperPath() ) as file:
        data = yaml.load(file, Loader=yaml.FullLoader)




    identities = []
    for user in users:
        identity = {
            "name": user.name,
            "clientPrivateKey": {
                "path":  str(user.private_key) 
            },
            "clientSignedCert": {
                "path":  str(user.certificate / "cert.pem") 
            }
        }
        identities.append(identity)

    # print() 
    data["organizations"][0]["identities"]["certificates"] = identities

    with open(GetCaliperPath(), "w") as file:
        yaml.dump(data, file)

    



def GetUserPaths() -> list[User]:
    """Returns a list of paths to the user's config files."""
    users : list[User] = []

    base = GetBasePath() / "organizations/peerOrganizations/org1.example.com/users/"

    for path in base.iterdir():

        private_key_dir = path / "msp" / "keystore"
        if private_key_dir.is_dir():
            private_key = next(private_key_dir.iterdir())


            user = User(
                name = path.name.split('@')[0],
                private_key = private_key,
                certificate = path / "msp" / "signcerts"
            )

            users.append(user)


    print(users)

    return users


def main():
    p = pathlib.Path()
    # print(GetBasePath())

    print(GetUserPaths())

    MakeYaml(GetUserPaths())

if __name__ == "__main__":
    main()
