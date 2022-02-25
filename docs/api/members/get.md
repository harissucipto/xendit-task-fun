# Show Members of an Organization

GET requests to `/orgs/<org-name>/members/` should return an array of members of an organization (with their login, avatar url, the numbers of number of people they're following), sorted in descending order by the number of followers.

**URL** : `orgs/<org-name>/members/`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

**Data constraints** : `{}`

## Success Responses

**Condition** : User can see one or more Members of an an Organization.

**Code** : `200 OK`

**Content** : In this example `orgs/xendit/members/`, the User can see members from organization registred against github (with their login, avatar url, the numbers of number of people they're following), sorted in descending order by the number of followers.

```json
[
  {
    "login": "bxcodec",
    "avatar_url": "https://avatars.githubusercontent.com/u/11002383?v=4",
    "followers": 637,
    "following": 54
  },
  {
    "login": "mkamadeus",
    "avatar_url": "https://avatars.githubusercontent.com/u/40513202?v=4",
    "followers": 126,
    "following": 150
  },
  {
    "login": "mychaelgo",
    "avatar_url": "https://avatars.githubusercontent.com/u/4651658?v=4",
    "followers": 81,
    "following": 27
  },
  {
    "login": "wildan3105",
    "avatar_url": "https://avatars.githubusercontent.com/u/7030099?v=4",
    "followers": 23,
    "following": 14
  },
  {
    "login": "fajarmf10",
    "avatar_url": "https://avatars.githubusercontent.com/u/30276682?v=4",
    "followers": 21,
    "following": 8
  },
  {
    "login": "Jafnee",
    "avatar_url": "https://avatars.githubusercontent.com/u/5748926?v=4",
    "followers": 17,
    "following": 45
  },
  {
    "login": "Pyakz",
    "avatar_url": "https://avatars.githubusercontent.com/u/49338297?v=4",
    "followers": 13,
    "following": 31
  },
  {
    "login": "acetdecastro",
    "avatar_url": "https://avatars.githubusercontent.com/u/13687580?v=4",
    "followers": 3,
    "following": 8
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
