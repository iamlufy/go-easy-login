### context: login

workFlow 

1. "authenticate"
   - triggered by: "user request to login"
   - input: "LoginCmd"
   - output: token which is the proof passing login or login fail detail message
   - assume: user's tenantId is existing and active 
   - flow   
     Authenticate
     1. check user if allow to login
        input: username
        for every login user should be
        - user is existing
        - active(not being locked)
          if not being allow to login
            outPut: detail reason
          else 
            outPut: true
     2. do verify password / sms code
        input: "login mode":password or smsCode,sourceCode,username 
          1. get encryptedCode by loginMode and username
          2. verify source and encryptedCode successfully 
        outPut: verify result 
     3. generate token which is the proof passing login
     
2. "add Login User"
   - triggered by: "request to add login user"
   - input:"AddLoginCmd"
   - output:"TBD"
   - assume: user's tenantId is existing and active 
   - flow 
     check
       - if user had existed => fail
       - username and mobile had existed =>fail 
 