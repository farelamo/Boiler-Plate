const vscode = require('vscode');
const fs  	 = require('fs');
const path	 = require('path');

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {
	console.log('Congratulations, your extension "mading" is now active!');

	const disposable = vscode.commands.registerCommand('mading.Option', function () {
		vscode.window.showQuickPick(['Golang', 'React Js']).then(option => {
            if (option) {
                createBoilerPlate(option, context.extensionPath);
            }else {
				vscode.window.showInformationMessage('Jirlah kok gajadi pake beb ? 😒 \n Jangan lupa mampir lagi ya 😏');
			}
        });
	});

	context.subscriptions.push(disposable);
}

function createBoilerPlate(option, extensionPath){
	const workspace = vscode.workspace.workspaceFolders ? vscode.workspace.workspaceFolders[0].uri.fsPath : undefined;
	if (!workspace) {
		vscode.window.showWarningMessage('Silahkan buka folder terlebih dahulu');
		return;
	}

	const boilerPLate = path.join(extensionPath, option);
	if (!fs.existsSync(boilerPLate)) {
		vscode.window.showErrorMessage(`Yahh, boilerplate '${option}' yang kamu pilih bermasalah, hubungi developer ya 🥺`);
		return;
	}

	copyBoilerPlate(option, boilerPLate, workspace)
}

function copyBoilerPlate(option, source, destination){
	if (!fs.existsSync(destination)) { //create destination folder if not exists
        fs.mkdirSync(destination);
    }

	if (fs.lstatSync(source).isDirectory()) {
		const files = fs.readdirSync(source);
		files.forEach((file) => {
			const sourcePath = path.join(source, file);
			const destPath   = path.join(destination, file)
	
			if (fs.lstatSync(sourcePath).isDirectory()) {
				if (file.startsWith('.')) {
					fs.copyFileSync(sourcePath, destPath);
				}else {
					copyBoilerPlate(option, sourcePath, destPath);
				}
			}else {
				fs.copyFileSync(sourcePath, destPath);
			}
		})
	}

	vscode.window.showInformationMessage(`Yeay! Boilerplate '${option}' kamu sudah siap. \n Kasih pendapat kamu ya 🥳`);
}

function deactivate() {}

module.exports = {
	activate,
	deactivate
}
