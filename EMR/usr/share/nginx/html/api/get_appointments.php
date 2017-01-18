<?php
require_once("../library/sqlconf.php");
require_once("../library/adodb/adodb.inc.php");
require_once("../library/adodb/drivers/adodb-mysql.inc.php");
require_once("../library/log.inc");

$pid = $_GET['id'];
$statement = "SELECT * FROM isemr_postcalendar_events WHERE pc_pid = " . $pid;
$recordset = mysql_query($statement);

$rows = array();
while($r = mysql_fetch_assoc($recordset)) {
    $rows['appointment'][] = $r;
}

if (sizeof($rows) <= 0) {
    $rows['appointment'] = array();
}


echo json_encode($rows);

