

## Keys 

### Item Keys
  When converted to its string form it will be:
   - `Key := <ITEM_TYPE>{COLLECTION_ID}{...ITEM_ID}`


#### SubKeys
  When converted to its string form it will be:
  `<SUB_ITEM_TYPE>{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUB_ITEM_ID}`

  - Examples:
      - Suggestion   := `<auth.Suggestion>    {COLLECTION_ID}{ITEM_TYPE} {...ITEM_ID}{SUGGESTION_ID}`
      - HiddenTxList := `<auth.HiddenTxList>  {COLLECTION_ID}{ITEM_TYPE} {...ITEM_ID}`


### Reference Keys
   Used to store references to items for case like a user having a role
   When converted to its string form it will be:
   - `<auth.Ref>{COLLECTION_ID}[{ITEM1_TYPE}{...ITEM1_ID}][{ITEM2_TYPE}{...ITEM2_ID}]`
   - `<auth.Ref>{COLLECTION_ID}[{ITEM2_TYPE}{...ITEM2_ID}][{ITEM1_TYPE}{...ITEM1_ID}]`

The Reference keys return a refrence object that includes the refKey and the serilized items that are referenced.



## Global

### Collections

#### Collection Types

- Allowed Item Types
- Allowed Reference Types
- Auth Types



### Users
