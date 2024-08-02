class Player {
    private name:string;
    private score:number;

    private isDrawing: boolean;

    constructor(name:string) {
        this.name = name;
        this.score = 0;
        this.isDrawing = false;
    }

    public getName(): string {
        return this.name;
    }

    public getScore(): number {
        return this.score;
    }

    public isDrawingPlayer(): boolean {
        return this.isDrawing;
    }

    public setDrawingStatus(status:boolean): void {
        this.isDrawing = status;
    }

    public incrementScore(points:number): void {
        this.score += points;
    }
}

export default Player