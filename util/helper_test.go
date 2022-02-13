package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenerateSignature(*testing.T) {
	key, _ := hex.DecodeString("1ae8ccd0e7985cc0b6203a55855a1034afc252980e970ca90e5202689f947ab9")
	iv, _ := hex.DecodeString("d58ce954203b7c9a9a9d467f59839249")

	plaintextBase64 := "ew0KICAicGF5bG9hZCIgOiAiQ0FFU3NBd0toUXNJQVJMc0NRcXZBZ2dDRWhGRGFISnZiV1ZEUkUwdFRXRmpMWGc0TmhqMXFJNk9CU0tPQWpDQ0FRb0NnZ0VCQUxrdTJ0VGp4R2FLQVFkTmRZWmM0QUQvZjI2Y3RqRG9YcGI0NlNCdEhmL0h4NWRuL1FQSVhPd2pJQWEwWFd1dXFycTJ6Sm8rWWovVDl0WFJrdDBrZ3E0MVkyem1nMmFLd01SY2dIbkRjdUQ1RnZwQ0RBNHR6MDAvT0hXakJINkFjWjFPM3hEVnVHVDA5TDRITlI2NkFXcEQ2b0hRV1F5aEhzMkNRTy84TnhqcDh3OHdQaldFbzF2SER1K0gzU2NvV2huWUp0R3Z1Y1ZPdVo3ZitUVlNQWmViM25PMVBkNythMHJGU21VMGRPUUFNVUdNRHdzdUovQ3g3RTAzdWdWalpCanNndWx5cll1VmRwUXJITTBWWUh4dlB6Q201WjFGWklHQyt3eEx5RUI1WUtqL0hLajFxc3RuK1lyYmFIWmovK0R5RVkzRi9hSnJjT1B0ZnpuMVRGTUNBd0VBQVNqOUlSS0FBblA2dEJMbWxxWHN3T3drNTA2eCthNlVWSjhQZUY1TklxSGpaVUk1dmNlSXM1WndqSkY4Rkd0V0xtcFVwcmRiNEpWV2JWTXc4RFBQSkJNajhDWElkRmNlc1VSWHlYWWthMnNWcEpka3hxalY3WDcyZXBUNlRZS1QwUVB2Q0hjeTBpenRHNGNDU1dqTVZiMGlYeDlQRjB4djlySS9ORktBSTJFVExFcnB4dEl1TlUrazRIVDhFVHRMdlhkWGRWRGpjOG56cHkvMzh1empraE9DSzlNQXVQUXJ1eFNhbHlqUmplaThSaGpibG42YnJ4WnVvcmFEU0ZZdFZVbXBWcVE5bmxRdTJIZkdFOUtCeE9MNjQ5TVZOMzBMeDhJbVpmSTdOUFFLdjlpZG1jQnhsK1A1aVBra2VJNk9nQ3JUL1JiTXFYcTIzSnh6Q0dJczFsaU5VNnhPa1JBYXRBVUtyZ0lJQVJJUVN1aGlBQlA3aHBsbldIQ2pYelMyT0JqL3BJNk9CU0tPQWpDQ0FRb0NnZ0VCQUs2ckN5YzBaZGMycmRxTHR1dFNySm1XcENGa2Q4RmZCSWV4VFdxVmQ3M0ZURHBtSVdLelZJQmtsamE0WmZZdkFsN29qbWg5ZnVOTmFGQTVkVlpCY0o5dTRmelFBeFBQQmE3OEVOT0VLcTZxOG9NNkF1WS8wSHRBNFNmVmJ4bGdtMVJjMlVUTFFiWWZVSTBXcnBTanZBQ2JYUlhiSEtWd2VqalVMVzhZazVJYzJrbUNweXZuRkZoTmxtdFBpdEtyNHQzeEI3ZkxhbWJLQVhvVE5qMGJ2SHBnRk0wajZBOFh6VkdCTHpCbWtnalhpKy90MXVnT0pzZjhPTFI4M1NZcy9sbFFRWFdFV3drQ1FpcFpleGEyNU0zS3M4Y293dFdLSXZmN0hIeWpYamRSODZOK1ZtRFhsLy80SXZlYnl2MUpka1lXcUlvdStqRFJGQzVrZWlHemZZVUNBd0VBQVNqOUlSS0FBemdNSzkzajlacGJOTjQxSGF0WEozTFdnTFJheGNFMFVMZjZTSlRGbTVSbkQrSFJpQzJYWUNqaFVxZmpvclViWUo0OG02enNPN3VVSUsvUkI3ZzFUUTNWTWhTNGFVV09aUUlYUkFMWDN0NmNITnAzZHUrZDJTbFkxRWdnL0xDb2d1b1FGNUFYSHdYY24rVWxKQVlINTlIQ21pc1JyVElqdktva3VCMTd6VHVtK1FrWHJmazF5dy81UnpUYU9kb0hkQzMxYTNUL3YwNU5hbmNZSTR3WWRkdXhPOGJuYVhpWDdTK1k2Zk1KTldKT0NMNEdwYWpReGlEdkc4L2I3OExlYS8yNGprVTJNbmZtRkRsZ0F1VkM0dGxEZ2I3TVVaRUVSeFJmR0V6NTlEdURIdW01SnJzTmZTVWkrRnpxZjFyU211YXRDS2J0NDFaT3RCcy9zNDBPaTNBYW1rT05odXpFZ1hwZXdSVyttaHQ4ckdNLzFaSy9SbjR2c3B6Sm5qaVpKOEtoQmY4ckQ0RER2UWNSWFB0TmZ3TzZsWWIyL2t0bkhGOWtTRXcyNWF4OVZkYVFsTWhYUDRJRnNwVGx6ZVM0MnVsL3dFbi9wam1ST0hKYWVVSEwrZXMvZnlYSXVHclM0R1F2S2FldG9hNVE2UmR6dUF1U3NOc0Y3VHBkV0RZQzhSb2JDaEZoY21Ob2FYUmxZM1IxY21WZmJtRnRaUklHZURnMkxUWTBHaFlLREdOdmJYQmhibmxmYm1GdFpSSUdSMjl2WjJ4bEdoY0tDbTF2WkdWc1gyNWhiV1VTQ1VOb2NtOXRaVU5FVFJvWENnMXdiR0YwWm05eWJWOXVZVzFsRWdaTllXTlBVMWdhSVFvVWQybGtaWFpwYm1WZlkyUnRYM1psY25OcGIyNFNDVEV1TkM0NExqazJNaklJQ0FBUUFCZ0JJQUFTbXdFS21BRUtnUUVTRU1BVDJkWithMVlTamxKdy9ES2xaalVTRUlRdXJmZVA5MUhSbGV4bFBOY3Q5ZWtTRVBvVGQ1ZlBUMXQ4cjZsYmF4NVF0UzBTRUZlNlFLdnFTRnpndmxoVWEzcVJkY2dTRUl3dS8xSWZTbDVEaGEwZllpd3FHWWthRFhkcFpHVjJhVzVsWDNSbGMzUWlFSys3WVJGT2lJL1pKRE1haElJN2xuNUk0OXlWbXdZUUFSb1FQdTNya21NOERBSE9nbEZxSzE3QmpCZ0JJSURnMmNZRk1CVWFnQUpJRTQ2dzEyUkM2MjlpTng1cGZJVGZzY1BLVmtnbzV5MnU3dzBORnl1OGJqVkxSWE50MW9CeVJ0VitDL3l4V0JuallRaTVaWGNqTHFVcUNVT3M1MEEreHlxNXkxNytrRGxDQkdTOXNoUDJCa052UzBwSGU5OWdCcnh3cTlNSWxTZ2Y4ZUxrUWQ0RFcyVEc1SjRFTmpIZjQvOVlDNkhCT1EvamhNdlJyYk51VHRsZzNnR2pQcHAvUTJWZzhtODBJTlpYbER2SklqUnZyVksrbXdhNVNDdHNoU2dlQkl0SFY1YmNOSy8vK1l5UElEZFhwMVgrT204Z2tGQ1ZqS2E3QjdyN3hCQlNEblpjYzhyaWExWWFRdUo1Zk45UU8yMmZYUGhMRkpacTRzK0VVM3RsTUFFa1hwaGRhUWVBUHJZR04vcUtydFU1TktWZGFEdFd3eEtYdnAxVyIsDQogICJwcm92aWRlciIgOiAid2lkZXZpbmVfdGVzdCIsDQogICJjb250ZW50X2lkIjogInI3dGhFVTZJajlra014cUVnanVXZmc9PSIsDQogICJjb250ZW50X2tleV9zcGVjcyI6IFsNCiAgICB7ICJ0cmFja190eXBlIjogIlNEX0hEIiB9DQogIF0NCn0NCg=="
	message, _ := base64.StdEncoding.DecodeString(plaintextBase64)
	fmt.Println("Message:")
	fmt.Println(string(message))

	signature := GenerateSignature(key, iv, message)

	fmt.Println("signature:" + signature)

}
