$(document).ready(function(){

    $(".mobile_nav").click(function(){
        if( $(".mobile_desc").is(":visible") ) {
            $(".mobile_desc").slideUp();
        } else {
            $(".mobile_desc").slideDown();
        }
    });

    $(".open_footer").click(function() {
        if( $(".footer").is(":visible")) {
            $(".footer").slideUp();
        } else {
            $(".footer").slideDown();
        }
    })

    /*
    $(window).scroll(function() {
        if ($(window).scrollTop() == $(document).height() - $(window).height()) { 
            $(".footer").slideDown();
        } 
    });​
    */

})