<?php

$ignoreAuth = true;
include_once("../interface/globals.php");
require_once("$srcdir/sql.inc");
require_once("$srcdir/pid.inc");
require_once("$srcdir/patient.inc");
require_once("$srcdir/options.inc.php");


// Get JSON Data
$json_data = json_decode(stripslashes(file_get_contents("php://input")), true);


$result = sqlQueryNoLog("SELECT MAX(pid) + 1 AS pid FROM patient_data");

$newpid = 1;

if ($result['pid'] > 1) $newpid = $result['pid'];

$json_data['pubpid'] = $newpid;

echo updatePatientData($newpid, $json_data, true) - 1;
