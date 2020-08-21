
nhn.husky.EZCreator.createInIFrame({
    oAppRef: oEditors,
    elPlaceHolder: "ir1", //textarea에서 지정한 id와 일치해야 합니다. 
    //SmartEditor2Skin.html 파일이 존재하는 경로
    sSkinURI: "/reviews/pkg/assets/smarteditor2-2.10.0/js/smart_editor2/SmartEditor2Skin_ko_KR.html",  
    htParams : {
    // 툴바 사용 여부 (true:사용/ false:사용하지 않음)
    bUseToolbar : true,             
    // 입력창 크기 조절바 사용 여부 (true:사용/ false:사용하지 않음)
    bUseVerticalResizer : true,     
    // 모드 탭(Editor | HTML | TEXT) 사용 여부 (true:사용/ false:사용하지 않음)
    bUseModeChanger : true,         
    fOnBeforeUnload : function(){ // 실행 전 설정 내용
                      
         }
    },fOnAppLoad : function(){
          var contents = "기본으로 적용 적용될 내용을 입력합니다."
        //기존 저장된 내용의 text 내용을 에디터상에 뿌려주고자 할때 사용
        oEditors.getById["ir1"].exec("PASTE_HTML", [ contents ]);
    },
    fCreator : "createSEditor2"
});