---
date: "2026-09-15T10:29:06.012Z"
title: "2026년 37주차(9월 2주차)"
---

Kukui(레노버 크롬북 듀엣)를 갖고 논 경험을 주저리주저리 적는 문서입니다. 두서 없을 수 있습니다. 재밌게 노세요.

- ChromeOS에 Crostini 없이 Flatpak 설치하기: https://github.com/rycont/chromeos-flatpak
- ChromeOS에 Crostini 없이 Xournal++ 설치하기: https://github.com/rycont/xournalpp-chromeos-appimage
- Kukui pmOS에서 안되는 기능들 해금하기(디스플레이 출력, 카메라, Vulkan 1.3): https://github.com/rycont/mt8183-kukui-pmos

---

# 개발자 모드 켜기

모든 정보가 초기화된다.

1. 장비를 끈다
2. 전원버튼 + 볼륨 위 + 볼륨 아래를 10초간 누르고 뗀다
3. USB를 삽입하라는 안내가 나오면 볼륨 위 + 볼륨 아래를 딸깍 누른다
4. OS 검증을 비활성화한다.

이후 부팅할 때 마다 개발자모드 경고 메뉴가 나오는데, 거기서 '개발자 옵션' -> '내부 디스크에서 부팅'을 누를 수 있다. 이후 모든 데이터가 초기화되고, 개발자모드가 활성화 된다.

# 터미널 접속

Ctrl - Alt - F2(대충 f2 위치에 있는 키)를 누르면 터미널이 뜬다. `chronos` 사용자 이름으로 로그인할 수 있다. `sudo` 비밀번호는 별도로 없다.

# Rootfs 검증 비활성화

crosh 터미널(`Ctrl Alt T` 하고 `shell` 입력)에서 수행한다

```
sudo /usr/share/vboot/bin/make_dev_ssd.sh --remove_rootfs_verification --partition 2
sudo reboot
```

# SSH 설정

developer console에서 수행한다.

```
sudo /usr/libexec/debugd/helpers/dev_features_ssh
sudo passwd chronos
```

이후 타 터미널에서 `ssh chronos@크롬북_주소`로 접근할 수 있다.

```
sudo -i
dev_install --reinstall --only_bootstrap
```

```
sudo modprobe fuse
ls -l /dev/fuse
```

---

목표:
