# DOCS API

Where full URLs are provided in responses they will be rendered as if service
is running on 'http://localhost:8080/' or a test server.

## Open Endpoints

Open endpoints require no Authentication.

- [Show List Comments from an Orginzation](./comments/get.md) : `GET /orgs/<org-name>/comments`
- [Create a Comment to an Organization](./comments/post.md) :`POST /orgs/<org-name>/comments`
- [Delete Comments from an Organization](./comments/post.md) : `DELETE /orgs/<org-name>/comments`
- [Show List Members of an Organization](./members/get.md) : `GET /orgs/<org-name>/members`

### Comment related

Each endpoint manipulates or display comments against a given github
organization.

- [Show List Comments from an Orginzation](./comments/get.md) : `GET /orgs/<org-name>/comments`
- [Create a Comment to an Organization](./comments/post.md) :`POST /orgs/<org-name>/comments`
- [Delete Comments from an Organization](./comments/post.md) : `DELETE /orgs/<org-name>/comments`

### Member related

Each endpoint manipulates or display members against a given github
organization.

- [Show List Members of an Organization](./members/get.md) : `GET /orgs/<org-name>/members`
