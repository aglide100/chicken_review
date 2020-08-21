document.getElementById('write_date').value= new Date().toISOString().slice(0, -1);
    document.getElementById('date').value=new Date().toISOString().substring(0,10);

    function Submit() {
        //alert("sending....");
        document.storeForm.submit();
    }
