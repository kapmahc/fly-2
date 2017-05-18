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
  
  // ---------------
})
