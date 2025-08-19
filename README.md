# My_Project

This is an small API implementation for signup and signin,
During signup process API accepts unique useraname and password and stores it in the database ,and hash-es the password securely using cryptogrophy method. 
In signin process the API validates provided credentials, upon successful validation it creates an JWT token and for subsequent login request middleware authenticates the JWT token, if token is valid allows the client to acess else returns error .
