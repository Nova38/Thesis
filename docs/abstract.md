
## On chain role based property level access control

### Abstract

Currently, the available methods for access control in smart contracts are limited
to  ACLs (Access Control Lists), State Based Key-Level endorsement, and ABAC
(Attribute Based Access Control). These methods have some limitations in terms of
flexibility and cross-organizational collaboration. These methods are based on the
membership of a singing identity in a fabric MSP (Membership Service Provider) or
an attribute in an X.509 certificate. This means that the access control is based
on the certificate of the user and not on the properties of the world state. This
limits the flexibility of the admins of the collection to define role privileges
to off-chain properties that are controlled by the individual organizations.
In this paper, we propose a new method for access control in smart contracts that
is based on an on-chain role-based property level access control. This method defines
a collection object that is used to store the privileges of each role with respect to
the properties of a state object. The roles of each user are defined in a on-chain
object that maps the users signed identity to roles in various collections. This
allows for the same user to have different roles in different collections. The
privileges are also more fine grained then the current key-level endorsement as
they are defined on a property level and not on a key level. They also allow for
the privileges to be more precise then just read and write access. This
method allows for a more flexible and fine grained access control that is embedded
into the smart contract itself.



This paper introduces a novel approach to mange  collections of artifacts through smart contract access control, rooted in on-chain role-based property-level access control. This smart contract fascinates the lifecycle of these artifacts including allowing for the creation,  modification, removal, and historical auditing  of the artifacts through both direct and suggested  actions. This method introduces a collection object designed to store role privileges concerning state object properties. User roles are defined within an on-chain entity that maps users' signed identities to roles across different collections, enabling a single user to assume varying roles in distinct collections. Unlike existing key-level endorsement mechanisms, this approach offers finer-grained privileges by defining them on a per-property basis, not at the key level. The outcome is a more flexible and fine-grained access control system seamlessly integrated into the smart contract itself, empowering administrators to manage access with precision and adaptability across diverse organizational contexts.  This has the added benefit of allowing for the auditing of not only the history of the artifacts but also for the permissions granted to the users.   
