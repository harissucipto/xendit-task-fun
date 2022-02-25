# Show Comments from an Orginzation

GET requests to `/orgs/<org-name>/comments/` should return an array of all the comments that have been registered against the organization.

**URL** : `/orgs/<org-name>/comments`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

**Data constraints** : `{}`

## Success Responses

**Condition** : User can not see any Comments.

**Code** : `200 OK`

**Content example** : For the example above, when the <org-name> is sony `/orgs/sony/comments` the name organization is registred towards github but the comment is still empty `

**Content** : `{[]}`

### OR

**Condition** : User can see one or more Accounts.

**Code** : `200 OK`

**Content** : In this example `/orgs/xendit/comments`, the User can see three comments because if the organization is registred against github, and the comments already existed.

```json
[
  {
    "id": 92,
    "comment": "wonderful",
    "org": "xendit",
    "created_at": "2022-02-24T23:26:52.661703Z"
  },
  {
    "id": 93,
    "comment": "best",
    "org": "xendit",
    "created_at": "2022-02-24T23:27:25.392049Z"
  },
  {
    "id": 94,
    "comment": "awesome company",
    "org": "xendit",
    "created_at": "2022-02-24T23:27:37.492688Z"
  }
]
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

**Condition** : If the user not provided the organization name `/orgs//comments`.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{ "error": "<PROVIDED WHAT SERVER WANT NEEDED OR VALIDATION>" }
```
