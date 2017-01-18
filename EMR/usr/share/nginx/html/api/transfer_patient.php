<?php
include_once("../library/sql.inc");
include_once('../library/authentication/password_hashing.php');

$result_query = mysql_query("SELECT * from patient_data where pid = " . $_GET['pid']);
$result = mysql_fetch_assoc($result_query);
$username = $result['fname'] . $result['lname'] . substr($result['DOB'],2,2);
$password_hash_result = mysql_query("SELECT * from `users_secure` where username = '" . $username . "'");
$password_assoc = mysql_fetch_assoc($password_hash_result);
$password = omer_password_dehash($password_assoc['password'], $password_assoc['salt']);
$export = array();

$export['name']['given'] = $result['fname'];
$export['name']['family'] = $result['lname'];
$export['address']['streetAddressLine'] = $result['street'];
$export['address']['city'] = $result['city'];
$export['address']['postalCode'] = $result['postal_code'];
$export['address']['state'] = $result['state'];
$export['gender'] = $result['sex'];
$export['birthTime'] = $result['DOB'];
$export['ssn'] = $result['ss'];
$export['extension']['username'] = $username;
$export['extension']['password'] = $password;
$url = $_GET['URL'];
$cmd = "curl -i -X POST -d '" . json_encode($export) .  "' " .  $url;
echo $cmd;

header("Location: ../interface/patient_file/summary/demographics.php");

// api/client







