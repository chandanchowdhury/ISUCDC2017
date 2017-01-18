<?php
$ignoreAuth = true;
include_once("../interface/globals.php");
require_once("$srcdir/sql.inc");

$json_data = json_decode(stripslashes(file_get_contents("php://input")), true);

$statement = "INSERT INTO lists SET `type` = 'allergy',  pid=" . $json_data['pid'] . ",title = '". $json_data['title'] . "'";

if (array_key_exists('begdate', $json_data)) {
    $statement .= ",begdate =" . $json_data['begdate'];
}

if (array_key_exists('enddate', $json_data)) {
    $statement .= ",enddate" . $json_data['enddate'];
}

if (array_key_exists('diagnosis', $json_data)) {
    $statement .= ",diagnosis'" . $json_data['diagnosis'] . "'";
}

if (array_key_exists('reaction', $json_data)) {
    $statement .= ",reaction '" . $json_data['reaction'] . "'";
}

if (array_key_exists('comments', $json_data)) {
    $statement .= ",comments '" . $json_data['comments'] . "'";
}

if (array_key_exists('severity', $json_data)) {
    $statement .= ",severity '" . $json_data['severity'] . "'";
}

if (array_key_exists('occurrence', $json_data)) {
    $statement .= ",occurrence '" . $json_data['occurrence'] . "'";
}

echo sqlQueryNoLog($statement);
