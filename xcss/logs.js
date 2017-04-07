function cd(i){
      var now = new Date();
      now.setDate(now.getDate()-i);
      var dd = now.getDate();
      var mm = now.getMonth()+1;
      var yy = now.getFullYear();

      if(dd<10){
          dd='0'+dd
        }
      if(mm<10){
          mm='0'+mm
        }
        var now = dd+'.'+mm+'.'+yy;
      return now;
    }

window.onkeypress = function(e) {
        if ((e.which || e.keyCode) == 13) {
    	    window.history.go(-1);
        }
        if ((e.which || e.keyCode) == 36) {
            window.location.href = '/';
        }
    }
