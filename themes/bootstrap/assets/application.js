$(function(){
  // ---------------------
  $("a.nav-link").each(function() {
      if ($(this).attr("href") == window.location.pathname) {
          $(this).addClass("active");
      }
  });
  // ---------------------
  $("img.votes").click(function(e) {
      e.preventDefault();
      $.ajax({
          url: "/votes",
          type: 'POST',
          data: {
              point: $(this).data('point'),
              type: $(this).data('type'),
              id: $(this).data('id'),
          },
          success: function(rst) {
              alert(rst.message)
          }.bind(this)
      })
  });
  // ----------------
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
  // ---------------
})
