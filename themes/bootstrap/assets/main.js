$.ajaxSetup({
    beforeSend: function(xhr, settings) {
        if (settings.type == 'POST' || settings.type == 'PUT' || settings.type == 'DELETE') {
            xhr.setRequestHeader('X-CSRF-Token', $('meta[name="csrf-token"]').attr('content'));
        }
    }
});

Dropzone.options.myAwesomeDropzone = {
    init: function() {
        this.on("sending", function(file, xhr, formData) {            
            xhr.setRequestHeader('X-CSRF-Token', $('meta[name="csrf-token"]').attr('content'));
        });
    }
};

$(function() {

    // ---------------------
    $('p.markdown').each(function(i, n) {
        var txt = $(this).text();
        $(this).html(marked(txt));
    });
    // ---------------------
    $("a.nav-link").each(function() {
        if ($(this).attr("href") == window.location.pathname) {
            $(this).addClass("active");
        }
    });
    // ---------------------
    $('a[data-method="delete"]').click(function(e) {
        e.preventDefault();
        if (confirm($(this).data('confirm'))) {
            $.ajax({
                url: $(this).attr('href'),
                type: 'DELETE',
                success: function(rst) {
                    window.location.href = $(this).data('next');
                }.bind(this)
            })
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
    // ---------------------
    $('pre code').each(function(i, block) {
        hljs.highlightBlock(block);
    });
    // ---------------------
    new Clipboard('.clipboard');
    // ---------------------
});
