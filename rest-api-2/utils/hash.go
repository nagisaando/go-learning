package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// The cost factor (second argument) determines the number of times bcrypt will apply a hashing algorithm to the password.
	// it specifies how complex hashed password should be
	// higher cost is harder for attackers to get the password, but it takes longer time to generate and impact application
	// cost factor 12-14 is generally recommended for the balance between security and usability

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func IsValidPassword(password, hashedPassword string) bool {
	// when password is hashed using the above function, it includes some additional metadata with the hash (salt -> random data added to the password before hashing)
	// the combination of salt and cost is used to generate the final result
	// this compare function read the metadata from the stored hash and
	// takes plain password, adds the salt and hashes it with the same cost factor
	// and it compares if new hash matches the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
