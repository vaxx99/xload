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
            subm();
        }

      if ((e.which || e.keyCode) == 27) {
      var sw = document.getElementById('ws');
      var hi = document.getElementById('ih');
      var sc = document.getElementById('cs');
      var na = document.getElementById('an');
      var nb = document.getElementById('bn');
      var ds = document.getElementById('sd');
      var de = document.getElementById('ed');
      var dr = document.getElementById('rd');
      var ot = document.getElementById('to');
      var it = document.getElementById('ti');
      var du = document.getElementById('ud');
      sw.value="";
      hi.value="";
      sc.value="";
      na.value="";
      nb.value="";
      ds.value="";
      de.value="";
      dr.value="";
      ot.value="";
      it.value="";
      du.value="";
      }

      if ((e.which || e.keyCode) == 120) {
              window.close();
          }
      if ((e.ctrlKey && e.keyCode) == 36) {
      window.location.href = '/';
    }
    if (e.ctrlKey && e.keyCode == 8) {
    window.location.href = '/logs';
  }

    }

    function subm(){
      var xl = document.getElementById('ldr');
      xl.style.display = 'block';
      var f0 = document.getElementById('f0');
      f0.submit();
      }

function showHide(){
    var xl = document.getElementById('ldr');
    if (xl.style.display == 'none') {
        xl.style.display = 'block';
    } else {
        xl.style.display = 'none';
    }
}
