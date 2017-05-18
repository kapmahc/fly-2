$.ajaxSetup({
    beforeSend: function(xhr, settings) {
        if (settings.type == 'POST' || settings.type == 'PUT' || settings.type == 'DELETE') {
            xhr.setRequestHeader('X-CSRF-Token', $('meta[name="csrf-token"]').attr('content'));
        }
    }
});

$(function(){
  // ---------------------
  $("form[id^='fm-']").submit(function(e){
      e.preventDefault();
      var data = $(this).serialize();

      $(this).find('input[type="checkbox"]').each(function(i, e) {
        if (!e.checked) {
          delete data[e.id];
        }
      });

      var next = $(this).data('next');
      $.post(
        $(this).attr('action'),
        data
      ).done(function(rst){
        if(rst.message) {
          alert(rst.message)
        }
        window.location.href = next;
      }).fail(function(req){
        alert(req.responseText || req.statusText);
      });
  });
  // ----------
})
