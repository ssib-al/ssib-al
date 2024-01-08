export const ShareKakao = (link: string) => {
  if (window.Kakao) {
    if (!window.Kakao.isInitialized()) {
      window.Kakao.init('05c6bb313b80d8e62c62c2fa11d963f0');
    }

    window.Kakao.Share.sendDefault({
      objectType: 'text',
      text: 'ssib.al 링크단축 - 공유된 링크를 확인해보세요',
      link: {
        mobileWebUrl: link,
        webUrl: link,
      },
    });
  }
};

export const ShareTwitter = (link: string) => {
  window.open(
    `https://twitter.com/intent/tweet?text=ssib.al 링크단축 - 공유된 링크를 확인해보세요&url=${link}`,
    '_blank',
  );
};
