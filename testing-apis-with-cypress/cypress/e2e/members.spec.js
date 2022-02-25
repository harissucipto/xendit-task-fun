describe("Show list Members API testing [GET /orgs/<org-name>/comments]", () => {
  it("have registered org name", () => {
    cy.request("/orgs/xendit/members/", { timeout: 10000 }).as(
      "membersRequest"
    );
    cy.get("@membersRequest").then(({ status, body }) => {
      expect(status).to.eq(200);
      expect(body).to.have.length.greaterThan(1);
      expect(body[0]).to.have.property("login");
      expect(body[0]).to.have.property("avatar_url");
      expect(body[0]).to.have.property("followers");
      expect(body[0]).to.have.property("following");
    });
  });

  let buildProp = (orgName) => [
    {
      failOnStatusCode: false,
      url: `/orgs/${orgName ? orgName : ""}/members`.trim(""),
      method: "GET",
    },
    { timeout: 10000 },
  ];

  it("cannot show  list when have not registred org name", () => {
    cy.request(...buildProp("234324")).as("membersRequest");
    cy.get("@membersRequest").then(({ status, body }) => {
      expect(status).to.eq(404);
    });
  });

  it("cannot show list when have empty org name ", () => {
    cy.request(...buildProp()).as("membersRequest");
    cy.get("@membersRequest").then(({ status, body }) => {
      expect(status).to.eq(400);
    });
  });

  it("ssorted list in descending order by the number of followers", () => {
    cy.request(...buildProp("xendit")).as("membersRequest");
    cy.get("@membersRequest").then(({ status, body }) => {
      expect(status).to.eq(200);
      expect(body).to.have.length.greaterThan(1);
      expect(body[0]).to.have.property("followers");
      expect(body[0]).to.have.property("following");
      expect(body[0]).to.have.property("login");
      expect(body[0]).to.have.property("avatar_url");
      expect(body[0].followers).to.be.greaterThan(body[1].followers);
    });
  });
});
