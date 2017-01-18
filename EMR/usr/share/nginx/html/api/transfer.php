<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <title>ISEmr Transfer</title>
</head>
<body>
    <h1><img src="../interface/pic/logo.gif"/></h1>
    <form method="get" action="transfer_patient.php">
        <label>P Id</label>
        <input type="text" name="pid" value="<?php echo $_GET['pid'] ?>" />
        <br />
        <label>Transfer Location URL</label>
        <input type="text" name="URL" />
        <br />
        <input type="submit" />
    </form>
</body>
</html>