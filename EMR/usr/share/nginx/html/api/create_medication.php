<?php
$ignoreAuth = true;
include_once("../interface/globals.php");
require_once("$srcdir/sql.inc");

$json_data = json_decode(stripslashes(file_get_contents("php://input")), true);

$statement = "INSERT INTO lists SET `type` = 'medication',  pid=" . $json_data['pid'] . " ,title = '". $json_data['title'] . "'";

if (array_key_exists('begdate', $json_data)) {
    $statement .= ",begdate =" . $json_data['begdate'];
}

if (array_key_exists('enddate', $json_data)) {
    $statement .= ",enddate" . $json_data['enddate'];
}

if (array_key_exists('diagnosis', $json_data)) {
    ",diagnosis'" . $json_data['diagnosis'] . "'";
}

echo sqlQueryNoLog($statement);