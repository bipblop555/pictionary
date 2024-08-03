import Player from "./Player";

class DrawingApp {
    private canvas: HTMLCanvasElement;
    private context: CanvasRenderingContext2D;
    private paint: boolean;

    private players: Player[];
    private currentPlayerIndex: number;
    private currentWord: string;

    private clickX: number[] = [];
    private clickY: number[] = [];
    private clickDrag: boolean[] = [];


    constructor(players: string[]){
        let canvas = document.getElementById('canvas') as HTMLCanvasElement;
        let context = canvas.getContext('2d') as CanvasRenderingContext2D;


        let paint:boolean = false;

        context!.lineCap = 'round';
        context!.lineJoin = "round";
        context!.strokeStyle = "black";
        context!.lineWidth = 1;

        this.canvas = canvas;
        this.context = context;
        this.paint = paint;

        this.players = players.map(name => new Player(name))
        this.currentPlayerIndex = 0;
        this.currentWord = "";

        this.redraw();
        this.createUserEvents();
        this.startGame();
    }

    private startGame(){
        this.currentWord = this.getRandomWord();
        this.players[this.currentPlayerIndex].setDrawingStatus(true);
        this.updateUI();
    }

    private getRandomWord(): string {
        const words = ['apple', "banana"];
        return words[Math.floor(Math.random() * words.length)]
    }

    private nextTurn() {
        this.players[this.currentPlayerIndex].setDrawingStatus(false);
        this.currentPlayerIndex = (this.currentPlayerIndex +1) % this.players.length;
        this.players[this.currentPlayerIndex].setDrawingStatus(true);
        this.currentWord = this.getRandomWord();
        this.clearCanvas();
        this.updateUI();
    }

    private updateUI(){
        document.getElementById('currentPlayer')!.textContent = `Current Player : ${this.players[this.currentPlayerIndex].getName()}`;
        document.getElementById('currentWord')!.textContent = `Current Word : ${this.currentWord}`;
    }

    private createUserEvents() {
        let canvas = this.canvas;

        canvas.addEventListener("mousedown", this.pressEventHandler);
        canvas.addEventListener("mousemove", this.dragEventHandler);
        canvas.addEventListener("mouseup", this.releaseEventHandler);
        canvas.addEventListener("mouseout", this.cancelEventHandler);

        canvas.addEventListener("touchstart", this.pressEventHandler);
        canvas.addEventListener("touchmove", this.dragEventHandler);
        canvas.addEventListener("touchend", this.releaseEventHandler);
        canvas.addEventListener("touchcancel", this.cancelEventHandler);

        document.getElementById('clear')!.addEventListener('click', this.clearEventHandler);
        document.getElementById('guess')!.addEventListener('submit', this.guessEventHandler);

        let docs = document.getElementsByClassName('colors') as HTMLCollection;
        for(let colors of docs){
            colors.addEventListener('click', this.getColorValue);
        }
    }


    private redraw() {
        let clickX = this.clickX;
        let clickY = this.clickY;

        let context = this.context;
        let clickDrag = this.clickDrag;

        for(let i = 0; i < clickX.length; i ++){
            context.beginPath();
            if(clickDrag[i] && i){
                context.moveTo(clickX[i - 1], clickY[i - 1]);
            } else {
                context.moveTo(clickX[i] - 1, clickY[i]);
            }

            context.lineTo(clickX[i], clickY[i])
            context.stroke()
        }
        context.closePath();
    }
    
    private addClick(x: number, y:number, dragging:boolean){
        this.clickX.push(x);
        this.clickY.push(y);
        this.clickDrag.push(dragging);
    }

    private clearCanvas(){
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.clickX = [];
        this.clickY = [];
        this.clickDrag = [];
    }

    private clearEventHandler = () => {
        this.clearCanvas();
    }

    private releaseEventHandler = () => {
        this.paint = false;
        this.redraw()
    }

    private cancelEventHandler = () => {
        this.paint = false;
    }

    private pressEventHandler = (e: MouseEvent | TouchEvent) => {
        if (!this.players[this.currentPlayerIndex].isDrawingPlayer()) return;
        let mouseX = (e as TouchEvent).changedTouches ?
                     (e as TouchEvent).changedTouches[0].pageX :
                     (e as MouseEvent).pageX;

        let mouseY = (e as TouchEvent).changedTouches ?
                     (e as TouchEvent).changedTouches[0].pageY :
                     (e as MouseEvent).pageY;

        mouseX -= this.canvas.offsetLeft;
        mouseY -= this.canvas.offsetTop;

        this.paint = true;
        this.addClick(mouseX, mouseY, false);
        this.redraw()
    }

    private dragEventHandler = (e: MouseEvent | TouchEvent) => {
        if (!this.players[this.currentPlayerIndex].isDrawingPlayer()) return;
        let mouseX = (e as TouchEvent).changedTouches ?
                     (e as TouchEvent).changedTouches[0].pageX : 
                     (e as MouseEvent).pageX;

        let mouseY = (e as TouchEvent).changedTouches ? 
                     (e as TouchEvent).changedTouches[0].pageY : 
                     (e as MouseEvent).pageY;
        
        mouseX -= this.canvas.offsetLeft;
        mouseY -= this.canvas.offsetTop;

        if (this.paint) {
            this.addClick(mouseX, mouseY, true);
            this.redraw();
        }

        e.preventDefault();
    }

    private guessEventHandler = (e: Event) => {
        e.preventDefault();

        const guessInput = document.getElementById('guessInput') as HTMLInputElement;
        const guess = guessInput.value.trim().toLowerCase();

        if (guess === this.currentWord.toLowerCase()) {
            alert(`${this.players[this.currentPlayerIndex].getName()} guessed the word`)
            this.players[this.currentPlayerIndex].incrementScore(1);
        } else {
            alert("Incorrect");
        }

        guessInput.value = "";
        this.nextTurn();
    }

    private getColorValue = (e: Event) => {
        e.preventDefault();

        const target = e.target as HTMLButtonElement;
        const element = target.className.split(' ');

        this.context.strokeStyle = element[1];
    }
}

export default DrawingApp