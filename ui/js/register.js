
        var message = "Led Zeppelin- Stairway to Heaven"; 
    
        // 생성된 해쉬 값을 사용하면 평문 사용을 피할 수 있다.
        var hashedPassword = CryptoJS.SHA256("1234").toString() ; 
    
        var encrypt = CryptoJS.AES.encrypt(message, hashedPassword);
        var decrypted = CryptoJS.AES.decrypt(encrypt, hashedPassword );
    
        // 암호화 이전의 문자열은 toString 함수를 사용하여 추출할 수 있다.
        var text = decrypted.toString(CryptoJS.enc.Utf8);