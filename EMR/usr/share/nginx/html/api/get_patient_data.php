<?php
require_once("../library/sqlconf.php");
require_once("../library/adodb/adodb.inc.php");
require_once("../library/adodb/drivers/adodb-mysql.inc.php");
require_once("../library/log.inc");

$statement = "SELECT * FROM patient_data";
$recordset = mysql_query($statement);

$rows = array();
while($r = mysql_fetch_assoc($recordset)) {
    $rows['patients'][] = $r;
}

if (sizeof($rows) <= 0) {
    $rows['patients'] = array();
}



echo json_encode($rows);

