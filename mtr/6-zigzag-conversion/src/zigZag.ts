var convert = function convert(s: string, numRows: number): string {
    if (numRows == 1) {return s}

    var travelUp = false; // if not traveling down, then we're traveling up
    var grid:string[][] = [];
    var yIndex = 0;
    var ret:string = "";

    // init grid
    for (var i = 0; i < numRows; i++) {
        grid.push([]);
    }

    // build the grid
    for (var i = 0; i < s.length; i++) {
        if (!travelUp) {
            // traveling down means adding a new row
            grid[yIndex].push(s[i]);
            
            if (yIndex == numRows - 1) {
                // at the bottom
                travelUp = true;
                continue;
            }

            yIndex++;
        }
        else {
            // traveling up which means adding to an existing row
            yIndex--;
            grid[yIndex].push(s[i]);

            if (yIndex == 0) {
                // at the top
                travelUp = false;
                yIndex++;
            }
        }
    }

    // build the string
    grid.forEach((row:string[]) => {
        ret += row.join("");
    });

    return ret;
};
export {convert};