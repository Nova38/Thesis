
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
is based on an on-chain role based property level access control. This method defines 
a collection object that is used to store the privileges of each role with respect to 
the properties of a state object. The roles of each user are defined in a on-chain 
object that maps the users signed identity to roles in various collections. This 
allows for the same user to have different roles in different collections. The 
privileges are also more fine grained then the current key-level endorsement as 
they are defined on a property level and not on a key level. They also allow for 
the privileges to be more precise then just read and write access. This
method allows for a more flexible and fine grained access control that is embedded
into the smart contract itself.   

