PROBLEM

Alice and Bob want to exchange messages. Eve should not be able to read the messages.

SOLUTION

Messages have to be transferred encrypted. We'll use simple shift cipher to do this. It encrypts by adding a key to each plaintext message byte. It decrypts by substracting a key from each ciphertext message byte.

Start implementing the solution by writing tests first (TDD).