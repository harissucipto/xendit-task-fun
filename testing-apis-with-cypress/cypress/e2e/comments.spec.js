let buildProp = (orgName) => [
  {
    failOnStatusCode: false,
    url: `/orgs/${orgName ? orgName : ""}/comments`.trim(""),
    method: "GET",
  },
  { timeout: 10000 },
];

describe(`Show List Comments from an Orginzation: GET /orgs/<org-name>/comments`, () => {
  it("should return a list of comments", () => {
    // create a comment request
    cy.request(
      {
        method: "POST",
        url: `/orgs/xendit/comments`,
        body: {
          comment: "This is a test comment",
        },
        failOnStatusCode: true,
      },
      {
        timeout: 10000,
      }
    );
    cy.request(...buildProp("xendit")).as("commentRequest");
    // check if the comment is in the list
    cy.get("@commentRequest").then((response) => {
      const commentId = response.body[response.body.length - 1].id;
      cy.request(...buildProp("xendit")).as("commentRequest");
      cy.get("@commentRequest").then((response) => {
        const comment = response.body.find(
          (comment) => comment.id === commentId
        );
        expect(comment).to.exist;
        expect(comment.comment).to.equal("This is a test comment");
      });
    });
  });

  it("valid organization but dont have any comment", () => {
    cy.request(...buildProp("sony")).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      expect(response.body).to.be.empty;
    });
  });

  it("organization not registred", () => {
    cy.request(...buildProp("hariskompany")).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      expect(response.status).to.equal(404);
    });
  });

  it("organization name is empty", () => {
    cy.request(...buildProp()).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      expect(response.status).to.equal(400);
    });
  });
});

describe(`[Create a Comment to an Organization](./comments/post.md) :POST /orgs/<org-name>/comments`, () => {
  it("should create a comment", () => {
    cy.request(
      {
        method: "POST",
        url: `/orgs/xendit/comments`,
        body: {
          comment: "This is a test comment",
        },
        failOnStatusCode: true,
      },
      {
        timeout: 10000,
      }
    );
    cy.request(...buildProp("xendit")).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      const commentId = response.body[response.body.length - 1].id;
      cy.request(...buildProp("xendit")).as("commentRequest");
      cy.get("@commentRequest").then((response) => {
        const comment = response.body.find(
          (comment) => comment.id === commentId
        );
        expect(comment).to.exist;
        expect(comment.comment).to.equal("This is a test comment");
      });
    });
  });

  it("cannot create a comment with empty comment or empty organization name", () => {
    cy.request(
      {
        method: "POST",
        url: `/orgs//comments`,
        body: {
          comment: "",
        },
        failOnStatusCode: false,
      },
      {
        timeout: 10000,
      }
    ).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      expect(response.status).to.equal(400);
    });
  });

  it("cannot create a comment with not registred organization name", () => {
    cy.request(
      {
        method: "POST",
        url: `/orgs/hariskompany/comments`,
        body: {
          comment: "This is a test comment",
        },
        failOnStatusCode: false,
      },
      {
        timeout: 10000,
      }
    ).as("commentRequest");
    cy.get("@commentRequest").then((response) => {
      expect(response.status).to.equal(404);
    });
  });
});

describe(`[Delete Comments from an Organization] DELETE /orgs/<org-name>/comments`, () => {
  it("can delete comment with name of Organization registred", () => {
    cy.request(
      {
        method: "POST",
        url: `/orgs/xendit/comments`,
        body: {
          comment: "This is a test comment",
        },
        failOnStatusCode: true,
      },
      {
        timeout: 10000,
      }
    );
    // deleted comment
    cy.request(
      {
        method: "DELETE",
        url: `/orgs/xendit/comments`,
        failOnStatusCode: true,
      },
      {
        timeout: 10000,
      }
    ).as("commentDeleteRequest");
    cy.get("@commentDeleteRequest").then((response) => {
      expect(response.status).to.equal(204);
    });
  });

  it("cannot delete comment with name of Organization not registred", () => {
    cy.request(
      {
        method: "DELETE",
        url: `/orgs/hariskompany/comments`,
        failOnStatusCode: false,
      },
      {
        timeout: 10000,
      }
    ).as("commentDeleteRequest");
    cy.get("@commentDeleteRequest").then((response) => {
      expect(response.status).to.equal(404);
    });
  });

  it("cannot delete comment with empty name of Organization", () => {
    cy.request(
      {
        method: "DELETE",
        url: `/orgs//comments`,
        failOnStatusCode: false,
      },
      {
        timeout: 10000,
      }
    ).as("commentDeleteRequest");
    cy.get("@commentDeleteRequest").then((response) => {
      expect(response.status).to.equal(400);
    });
  });
});
