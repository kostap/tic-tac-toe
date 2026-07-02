class TicTacToe {
    constructor() {
        this.board = document.getElementById('board');
        this.status = document.getElementById('status');
        this.newGameBtn = document.getElementById('newGame');
        this.cells = document.querySelectorAll('.cell');

        this.init();
    }

    init() {
        this.cells.forEach(cell => {
            cell.addEventListener('click', (e) => this.handleCellClick(e));
        });

        this.newGameBtn.addEventListener('click', () => this.startNewGame());
        this.startNewGame();
    }

    async startNewGame() {
        try {
            const response = await fetch('/api/new-game');
            const data = await response.json();
            this.updateBoard(data);
        } catch (error) {
            console.error('Error:', error);
        }
    }

    async handleCellClick(e) {
        const cell = e.target;
        const row = parseInt(cell.dataset.row);
        const col = parseInt(cell.dataset.col);

        if (cell.textContent) return;

        try {
            const response = await fetch('/api/move', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ row, col }),
            });

            if (!response.ok) {
                const error = await response.text();
                alert(error);

                return;
            }

            const data = await response.json();

            this.updateBoard(data);
        } catch (error) {
            console.error('Error:', error);
        }
    }

    updateBoard(data) {
        this.cells.forEach(cell => {
            const row = parseInt(cell.dataset.row);
            const col = parseInt(cell.dataset.col);

            cell.textContent = data.board[row][col];

            if (data.board[row][col] === 'X') {
                cell.className = 'cell x';
            } else if (data.board[row][col] === 'O') {
                cell.className = 'cell o';
            } else {
                cell.className = 'cell';
            }
        });

        this.status.textContent = data.message;

        if (data.gameOver) {
            this.cells.forEach(cell => {
                cell.style.pointerEvents = 'none';
            });
        } else {
            this.cells.forEach(cell => {
                cell.style.pointerEvents = 'auto';
            });
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    new TicTacToe();
});