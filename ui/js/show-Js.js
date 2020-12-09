
//var swiper;
var lat = document.getElementById("lat").value;
var lng = document.getElementById("lng").value;

if ((lat == "") || (lng == "") || (lat == undefined) || (lng == undefined)){
    // 혹시라도 지도가 표시될 경우
    lat = 33.450701;
    lng = 126.570667;
}

var mapContainer = document.getElementById('map'), // 지도를 표시할 div 
    mapOption = { 
        
        center: new kakao.maps.LatLng(lat, lng), // 지도의 중심좌표
        level: 3 // 지도의 확대 레벨
    };


var map = new kakao.maps.Map(mapContainer, mapOption); // 지도를 생성합니다

// 마커가 표시될 위치입니다 
var markerPosition  = new kakao.maps.LatLng(lat, lng); 

// 마커를 생성합니다
var marker = new kakao.maps.Marker({
    position: markerPosition
});

// 마커가 지도 위에 표시되도록 설정합니다
marker.setMap(map);

// 아래 코드는 지도 위의 마커를 제거하는 코드입니다
// marker.setMap(null);    


function deleteReview() {
    var check = confirm("delete review?");
    if(check == true) {
        var para = document.location.href.split("reviews/");
        console.log(para)
        var URI = "delete/";
        const str = [para[0], URI, para[1]].join('');
        console.log(str)
        location.href = str;
    }
}

var swiper = new Swiper('.swiper-container', {
    slidesPerView: 3,
    spaceBetween: 30,
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
});