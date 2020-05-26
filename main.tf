provider "cloudknox" {
    service_account_id = "SAU0HMGINBOQNQ4C"
    access_key = "SAAKBQHF2OGOBD5Q"
    secret_key = "BCW9qDmvYRQeluTo"
    shared_credentials_file = "C:\\Users\\Saketh Kollu\\.cnx\\credentials"
}

resource "cloudknox_policy" "new-policy" {
    address = "1"
}