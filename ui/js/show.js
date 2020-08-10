
function initialize() {
    var myLatlng = new google.maps.LatLng(37.556911,126.918694);

    var myOptions = {
        zoom: 15,
        center: myLatlng,
        mapTypeId: google.maps.MapTypeId.TERRAIN
    }
    var map = new google.maps.Map(document.getElementById("map_canvas"), myOptions);

    var marker = new google.maps.Marker( {
        position: myLatlng, 
        map: map, 
        title:"Seoul" 
    });   
    var infowindow = new google.maps.InfoWindow( {
        content: "Seoul gangnam"
    });
    infowindow.open(map,marker);
}
window.onload=function() {
    initialize();
    var mySwiper = new Swiper('.swiper-container', {
    // Optional parameters
    //direction: 'vertical',
    loop: true,

    // If we need pagination
    pagination: {
        el: '.swiper-pagination',
    },

    // Navigation arrows
    navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev',
    },

    // And if we need scrollbar
    /*
    scrollbar: {
        el: '.swiper-scrollbar',
    },
    */
    });

}
// Maps api end

function btnDelete() {
    var check = confirm("delete review?");
   
    if(check == true) {
        location.href = '/delete/{{.ID}}';
    }
    
}