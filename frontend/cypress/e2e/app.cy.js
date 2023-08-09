describe('Testing the app', () => {

  it('should log in, add a room, navigate to the room, send a message', () => {
    // Step 1: Login
    cy.visit('http://localhost:5173/login');
    cy.get('#username').type('lyes');
    cy.get('#password').type('hello');
    cy.get('#loginBtn').click();

    // Step 2: Navigate to the rooms page
    cy.url().should('eq', 'http://localhost:5173/');

    // Step 3: Adding rooms
    cy.get('#addroom').click();
    cy.get('#newRoomName').type('test');
    cy.get('#submitNewRoom').click();

    // Step 4: Navigate to the new room
    cy.url().should('include', '/test');


    // Step 5: Send a message

    cy.get('#newmessage').type('testin-gmessages');
    cy.get('#submitmessage').click();


  });

});
