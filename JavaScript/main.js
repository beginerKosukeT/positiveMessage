//ログアウト
function logout() {
    if (window.confirm("ログアウトしますか？")) {
        window.location = "/logout";
    }
}

// アカウント登録時にアイコンをクリックすると、選択状態にする
var previousIconNumber;
function iconClick(iconNumberSelected) {
    let inputIconNumber = document.getElementById('inputIconNumber');
    inputIconNumber.value = iconNumberSelected;
    if (previousIconNumber != null) {
        let previousIconSelected = document.getElementById('sg' + previousIconNumber);
        previousIconSelected.classList.remove('bg-primary');
    }
    let iconSelected = document.getElementById('sg' + iconNumberSelected);
    iconSelected.classList.add('bg-primary');
    previousIconNumber = iconNumberSelected;
}
// アカウント登録時、入力の不備をエラーとして表示する
function regisrationSubmitCheck() {
    if (document.getElementById('inputIconNumber').value == "" ||
        document.getElementById('userName').value == "" ||
        document.getElementById('emailAddress').value == "" ||
        (document.getElementById('password').value).length < 6) {
        alert("入力に不備があります。入力内容を確認してください。");
        return false;
    }
}

// 新規投稿時、入力の不備をエラーとして表示する
function postSubmitCheck() {
    if (document.getElementById('postTitle').value == "" ||
        document.getElementById('body').value == "") {
        alert("入力に不備があります。入力内容を確認してください。");
        return false;
    }
}


//メッセージを再生
var typeOfPosts;//マイページに表示される、保存済みの投稿、お気に入り、作成した投稿の内、どれを再生するかを管理する
var playCount = 0;//メッセージを再生または連続再生する時、先頭から何番目の投稿を再生するのかを管理する

function play(body) {//再生
    if (speechSynthesis.paused) {
        speechSynthesis.resume();
    } else if (speechSynthesis.speaking == false) {
        const uttr = new SpeechSynthesisUtterance(body);
        uttr.lang = 'ja-JP';
        uttr.rate = 1;
        speechSynthesis.speak(uttr);
        uttr.onend = function () {
            clickSkipForwardButton();
        };
        switch (typeOfPosts) {
            case 's':
                document.getElementById("nowPlayingInSavedPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'f':
                document.getElementById("nowPlayingInFavoritePosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'p':
                document.getElementById("nowPlayingInPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
        }
    }
}
function pause() {//一時停止
    speechSynthesis.pause();
}
function skip(body) {//次または前のメッセージを再生
    speechSynthesis.cancel();
    const uttr = new SpeechSynthesisUtterance(body);
    uttr.lang = 'ja-JP';
    uttr.rate = 1;
    speechSynthesis.speak(uttr);
    uttr.onend = function () {
        clickSkipForwardButton();
    };
    switch (typeOfPosts) {
        case 's':
            document.getElementById("nowPlayingInSavedPosts").innerHTML =
                document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
            break;
        case 'f':
            document.getElementById("nowPlayingInFavoritePosts").innerHTML =
                document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
            break;
        case 'p':
            document.getElementById("nowPlayingInPosts").innerHTML =
                document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
            break;
    }
}
function cancel() {//再生中のメッセージをキューから削除
    speechSynthesis.cancel();
    playCount = 0;
    switch (typeOfPosts) {
        case 's':
            document.getElementById("nowPlayingInSavedPosts").innerHTML = "";
            break;
        case 'f':
            document.getElementById("nowPlayingInFavoritePosts").innerHTML = "";
            break;
        case 'p':
            document.getElementById("nowPlayingInPosts").innerHTML = "";
            break;
    }
}

function clickSkipBackwardButton() {//マイページで前のメッセージを再生するために、各投稿に埋め込まれた非表示のボタンを押す
    if (document.getElementById(
        "hiddenSkipButton-" + typeOfPosts + (playCount - 1)
    ) != null) {
        playCount -= 1;
        switch (typeOfPosts) {
            case 's':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
            case 'f':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
            case 'p':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
        }
    }
}
function clickPlayButton(type) {//マイページでメッセージを再生するために、各投稿に埋め込まれた非表示のボタンを押す
    if (typeOfPosts != type) {
        typeOfPosts = type;
        speechSynthesis.cancel();
        playCount = 0;
        document.getElementById("nowPlayingInSavedPosts").innerHTML = "";
        document.getElementById("nowPlayingInFavoritePosts").innerHTML = "";
        document.getElementById("nowPlayingInPosts").innerHTML = "";
        switch (typeOfPosts) {
            case 's':
                document.getElementById("nowPlayingInSavedPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'f':
                document.getElementById("nowPlayingInFavoritePosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'p':
                document.getElementById("nowPlayingInPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
        }
    }
    document.getElementById(
        "hiddenPlayButton-" + typeOfPosts + playCount
    ).click();
}
function clickSkipForwardButton() {//マイページで次のメッセージを再生するために、各投稿に埋め込まれた非表示のボタンを押す
    if (document.getElementById(
        "hiddenSkipButton-" + typeOfPosts + (playCount + 1)
    ) == null) {
        playCount = 0;
        switch (typeOfPosts) {
            case 's':
                document.getElementById("nowPlayingInSavedPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'f':
                document.getElementById("nowPlayingInFavoritePosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
            case 'p':
                document.getElementById("nowPlayingInPosts").innerHTML =
                    document.getElementById("postOverview-" + typeOfPosts + playCount).innerHTML;
                break;
        }
    } else {
        playCount += 1;
        switch (typeOfPosts) {
            case 's':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
            case 'f':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
            case 'p':
                document.getElementById(
                    "hiddenSkipButton-" + typeOfPosts + playCount
                ).click();
                break;
        }
    }
}