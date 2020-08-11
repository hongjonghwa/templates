from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import UnexpectedAlertPresentException, TimeoutException
from bs4 import BeautifulSoup
import random
import sys
from time import sleep

options = webdriver.ChromeOptions()
# options = webdriver.FirefoxOptions()

# headless 옵션 설정
#options.add_argument('headless')
#options.add_argument("no-sandbox")

# 브라우저 윈도우 사이즈
#options.add_argument('window-size=1920x1080')
options.add_argument('window-size=1280x720')

# 사람처럼 보이게 하는 옵션들
options.add_argument("disable-gpu")   # 가속 사용 x
options.add_argument("lang=ko_KR")    # 가짜 플러그인 탑재
#options.add_argument('user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36')  # user-agent 이름 설정

# 드라이버 위치 경로 입력
driver = webdriver.Chrome('D:/templates/selenium/chromedriver.exe', options=options)
# driver = webdriver.Firefox(executable_path='D:\templates\selenium\geckodriver.exe', options=options)
driver.get('https://fafan.kr/')
driver.implicitly_wait(3)

# 로그인
user_id = input("아이디? ")
user_password = input("패스워드? ")

driver.find_element_by_id('txtLoginUserId').send_keys(user_id)
driver.find_element_by_id('txtLoginPasswd').send_keys(user_password)

driver.find_element_by_id('btnLogin').click()

# 페이지 이동
# driver.get(a[0].get_attribute('href'))

wait = WebDriverWait(driver, 10)
# wait.until(lambda driver: driver.current_url != "https://fafan.kr/default.aspx")
# driver.implicitly_wait(3)


"""
### 상자가 나올 때 ###
<div id="divEventFloating" style="display: inline; position: fixed; width: 60px; height: 60px; top: 50%; left: 50%; margin-left: 351px; margin-top: -30px;">
    <img src="/common/image/present.png" style="cursor:pointer;" onclick="document.getElementById('btnEventPoint').click();"
            alt="팝업광고가 아닙니다. 클릭하시면 보너스 포인트를 드립니다.">
    <img src="/common/image/btn_delete.gif" style="cursor:pointer;"
            onclick="if(confirm('확인을 클릭하시면\n하루동안 이벤트에 참여하지 않습니다')) { document.getElementById('btnCancel').click(); } return false;">
</div>
<div id='divEventFloating' style='display:none; position:fixed; text-aling:center; width:60px; height:60px; top:50%; left:50%; margin-left:351px; margin-top:-30px;'>
    <img src='/common/image/present.png' style='cursor:pointer;' onclick="document.getElementById('btnEventPoint').click();"
            alt='팝업광고가 아닙니다. 클릭하시면 보너스 포인트를 드립니다.' />
    <img src='/common/image/btn_delete.gif' style='cursor:pointer;'
            onclick="if(confirm('확인을 클릭하시면\n하루동안 이벤트에 참여하지 않습니다')) { document.getElementById('btnCancel').click(); } return false;">
</div>
<script>setTimeout("document.getElementById('divEventFloating').style.display = 'inline'",1000);</script>
### 고정 ###
<input type="submit" name="ctl00$btnEventPoint" value="" id="btnEventPoint" style="display: none" />
<input type="submit" name="ctl00$btnCancel" value="" id="btnCancel" style="display: none" />
<!-- 이벤트 //-->
"""



def contain_present(image_tags):
    for img in image_tags:
        # print(img.get('src'))
        if img.get('src') == '/common/image/present.png':
            return True
    return False

def open_present():
    a = driver.execute_script("document.getElementById('btnEventPoint').click()")
    # print (a)
    wait.until(EC.alert_is_present())
    alert = driver.switch_to.alert
    print(alert.text)
    alert.accept()
    driver.implicitly_wait(2)

present_count = 0
while True:
    try:
	    wait.until(EC.presence_of_element_located((By.CSS_SELECTOR,"#btnEventPoint")))
    except UnexpectedAlertPresentException as e:
        print('[!] Error: ' + str(e))
        #sys.exit()
        #driver.switch_to.alert.accept()
    except TimeoutException as t:
        print('[!] Timeout: ' + str(t))

    driver.implicitly_wait(3)
    sleep(10) # 충분히 길게하여 테스트 해보자!

    html = driver.page_source
    soup=BeautifulSoup(html, 'html.parser')

    anchors = soup.findAll('a')
    links = []
    for a in anchors:
        href = a.get('href')
        if href != None and ( href.startswith('/board/list.aspx?id=') or href.startswith('/board/view.aspx?id=') ):
            links.append(href)

    random_url = random.choice(links)
    print("random url is", random_url)

    if contain_present(soup.findAll('img')):
        open_present()
        present_count += 1
        print("Present no {}".format(present_count), flush=True)
    else:
        print("No Present!", flush=True)

    driver.get('https://fafan.kr'+random_url)

driver.get_screenshot_as_file('capture_fafan.png')    # 화면캡처
driver.quit() # driver 종료
