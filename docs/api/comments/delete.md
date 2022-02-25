# Delete Comments from an Organization

DELETE requests to /orgs/<org-name>/comments should soft delete all comments associated with a particular organization. We define a "soft delete" to mean that deleted items should not be returned in GET calls, but should remain in the database for emergency retrieval and audit purposes.

**URL** : `/orgs/<org-name>/comments`

**Method** : `DELETE`

**Auth required** : No

**Permissions required** : None

**Data** : `{}`

## Success Response

**Condition** : If the Organization exists.

**Code** : `204 NO CONTENT`

**Content** : `{}`

## Error Responses

**Condition** : If there was no Account available to delete.

**Code** : `404 NOT FOUND`

**Content** : `{}`

### Or

**Condition** : Authorized User is not Owner of Account at URL.

**Code** : `403 FORBIDDEN`

**Content** : `{}`

## Notes

- Will remove memberships for this Account for all Users that had access.
