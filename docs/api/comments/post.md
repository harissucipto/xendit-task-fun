# Create a Comment to an Organization

POST requests to `/orgs/<org-name>/comments` should allow the user to persist comments (in a MongoDB collection or Postgres table) against a given github organization.

**URL** : `/orgs/<org-name>/comments`

**Method** : `POST`

**Auth required** : No

**Permissions required** : None

**Data constraints**

Provide comment of Organization to be created.

```json
{
  "comment": "[string, required]"
}
```

**Data example** All fields must be sent.

```json
{
  "comment": "awesome company"
}
```

## Success Response

**Condition** : If everything is OK and an Organization registered against a given github organization.

**Code** : `201 CREATED`

**Content example**

```json
{
  "comment": "awesome company"
}
```

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
