document.getElementById('write_date').value= new Date().toISOString().slice(0, -1);
    document.getElementById('date').value=new Date().toISOString().substring(0,10);

    function Submit() {
        //alert("sending....");
        document.storeForm.submit();
    }



/**
         * Google Map API 주소의 callback 파라미터와 동일한 이름의 함수이다.
         * Google Map API에서 콜백으로 실행시킨다.
         */
        function initMap() {
            console.log('Map is initialized.');
 
            /**
             * 맵을 설정한다.
             * 1번째 파라미터 : 구글 맵을 표시할 위치. 여기서는 #google-map
             * 2번째 파라미터 : 맵 옵션.
             *      ㄴ zoom : 맵 확대 정도
             *      ㄴ center : 맵 중심 좌표 설정
             *              ㄴ lat : 위도 (latitude)
             *              ㄴ lng : 경도 (longitude)
             */
            var map = new google.maps.Map(document.getElementById('google-map'), {
                zoom: 12.5,
                center: {
                    lat: -34.397,
                    lng: 150.644
                }
            });
 
            /**
             * Google Geocoding. Google Map API에 포함되어 있다.
             */
            var geocoder = new google.maps.Geocoder();
 
            // submit 버튼 클릭 이벤트 실행
            document.getElementById('submit').addEventListener('click', function() {
                console.log('submit 버튼 클릭 이벤트 실행');
 
                // 여기서 실행
                geocodeAddress(geocoder, map);
            });
 
            /**
             * geocodeAddress
             * 
             * 입력한 주소로 맵의 좌표를 바꾼다.
             */
            function geocodeAddress(geocoder, resultMap) {
                console.log('geocodeAddress 함수 실행');
 
                // 주소 설정
                var address = document.getElementById('address').value;
 
                /**
                 * 입력받은 주소로 좌표에 맵 마커를 찍는다.
                 * 1번째 파라미터 : 주소 등 여러가지. 
                 *      ㄴ 참고 : https://developers.google.com/maps/documentation/javascript/geocoding#GeocodingRequests
                 * 
                 * 2번째 파라미터의 함수
                 *      ㄴ result : 결과값
                 *      ㄴ status : 상태. OK가 나오면 정상.
                 */
                geocoder.geocode({'address': address}, function(result, status) {
                    console.log(result);
                    console.log(status);
 
                    if (status === 'OK') {
                        resultMap.setCenter(result[0].geometry.location);
                        resultMap.setZoom(18);
                        var marker = new google.maps.Marker({
                            map: resultMap,
                            position: result[0].geometry.location
                        });
 
                        console.log('latitude : ' + marker.position.lat());
                        console.log('longitude : ' + marker.position.lng());
                    } else {
                        alert('지오코드가 다음의 이유로 성공하지 못했습니다 : ' + status);
                    }
                });
            }
        }