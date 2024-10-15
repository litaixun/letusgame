function startGame(game) {
    fetch(`/play/${game}`, {
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        const outputDiv = document.getElementById('game-output');
        outputDiv.innerHTML = `<h2>${data.message}</h2>`;
        outputDiv.style.display = 'block';

        // 清空之前的输入框
        document.getElementById('guess-input').style.display = 'none';
        document.getElementById('rps-input').style.display = 'none';
        document.getElementById('timer-input').style.display = 'none';

        if (game === 'guess') {
            document.getElementById('guess-input').style.display = 'block';
        } else if (game === 'rps') {
            document.getElementById('rps-input').style.display = 'block';
        } else if (game === 'timer') {
            document.getElementById('timer-input').style.display = 'block';
        }
    });
}

function submitGuess() {
    const guess = document.getElementById('user-guess').value;
    fetch('/play/guess', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ number: parseInt(guess) })
    })
    .then(response => response.json())
    .then(data => {
        displayOutput(data.message);
        if (data.message === "恭喜你，猜对了！") {
            triggerCelebration();
        }
    });
}

function submitRPS(choice) {
    fetch('/play/rps', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ choice: choice })
    })
    .then(response => response.json())
    .then(data => {
        displayOutput(`计算机选择了 ${data.computer}。 ${data.result}`);
        if (data.result === "你赢了！") {
            triggerCelebration();
        }
    });
}

function submitTimer() {
    fetch('/play/timer', {
        method: 'POST',
    })
    .then(response => response.json())
    .then(data => {
        displayOutput(data.message);
        if (data.message === "成功在 5 秒内提交！") {
            triggerCelebration();
        }
    });
}

function displayOutput(message) {
    const outputDiv = document.getElementById('game-output');
    outputDiv.innerHTML += `<p>${message}</p>`;
}

function triggerCelebration() {
    const celebrationDiv = document.getElementById('celebration');
    celebrationDiv.style.display = 'block';

    // 动态庆祝彩条和烟花
    createFireworks(5);
}

function createFireworks(num) {
    const fireworksContainer = document.createElement('div');
    fireworksContainer.classList.add('fireworks');

    for (let i = 0; i < num; i++) {
        const firework = document.createElement('div');
        firework.classList.add('firework');
        firework.style.width = `${Math.random() * 20 + 10}px`;
        firework.style.height = firework.style.width;
        firework.style.left = `${Math.random() * 100}vw`;
        firework.style.bottom = `0`;
        firework.style.backgroundColor = getRandomColor();
        fireworksContainer.appendChild(firework);

        setTimeout(() => {
            firework.remove();
        }, 1000);
    }

    celebrationDiv.appendChild(fireworksContainer);
    setTimeout(() => {
        fireworksContainer.remove();
    }, 2000);
}

function getRandomColor() {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}

