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

**Condition** : If the name organization is not registred towards github.

**Code** : `404 NOT FOUND`

**Content** :

```json
{
  "error": "org not found"
}
```

### Or

**Condition** : If the github or github token is not valid.

**Code** : `404 NOT FOUND`

**Content** :

```json
{ "error": "github api error" }
```

### Or

**Condition** : If the user not provided the organization name `/orgs//comments` or the content contstraint is not valid.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{ "error": "<PROVIDED WHAT SERVER WANT NEEDED OR VALIDATION>" }
```
