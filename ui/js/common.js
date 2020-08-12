$(document).ready(function(){
    $(".mobile_nav").click(function(){
        if( $(".mobile_desc").is(":visible") ) {
            $(".mobile_desc").slideUp();
        } else {
            $(".mobile_desc").slideDown();
        }
    });
})