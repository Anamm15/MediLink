package constants

const EmailOTPTemplate = `
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background: #ffffff; padding: 30px; border-radius: 8px; box-shadow: 0 0 10px rgba(0,0,0,0.1); }
        .header { text-align: center; margin-bottom: 20px; }
        .header h1 { color: #333; }
        .otp-box { background-color: #e0f7fa; color: #006064; font-size: 32px; font-weight: bold; text-align: center; padding: 15px; margin: 20px 0; border-radius: 5px; letter-spacing: 5px; }
        .footer { font-size: 12px; color: #888; text-align: center; margin-top: 30px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Verifikasi MediLink</h1>
        </div>
        <p>Halo <strong>%s</strong>,</p>
        <p>Anda sedang melakukan proses verifikasi akun. Gunakan kode OTP di bawah ini untuk melanjutkan. Kode ini berlaku selama 3 menit.</p>
        
        <div class="otp-box">%s</div>
        
        <p>Jika Anda tidak meminta kode ini, abaikan email ini. Jangan berikan kode ini kepada siapapun.</p>
        
        <div class="footer">
            &copy; 2024 MediLink Healthcare System. All rights reserved.
        </div>
    </div>
</body>
</html>
`
