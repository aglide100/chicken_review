$(document).ready(function(){


    $(".search_box").click(function(){
        
            if( $(".map_wrap").is(":visible") ) {
                $(".map_wrap").slideUp();
            } else {
                $('html').scrollTop(0);
                $(".map_wrap").slideDown('fast', function() {
                    reloadLayout();
                });
            }
            
        /*
        $.when( ShowMap() ).done(function() {
            setTimeout(function() {
                reloadLayout();
              }, 1000);
            
        });
        */
    });
    
    

    $(".close_kakaomap").click(function(){
        $(".map_wrap").slideUp();
    })
})