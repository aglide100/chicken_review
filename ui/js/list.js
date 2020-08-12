function Search() {
    document.searchForm.submit();
}

$(document).ready(function(){
    $(".Search-button").click(function(){
        if( $(".searchForm").is(":visible") ) {
            $(".searchForm").slideUp();
        } else {
            $(".searchForm").slideDown();
        }
    });
})