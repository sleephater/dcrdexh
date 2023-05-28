package locales

var DeDE = map[string]string{
	"Language":                       "de-DE",
	"Markets":                        "Märkte",
	"Wallets":                        "Wallets",
	"Notifications":                  "Benachrichtigungen",
	"Recent Activity":                "Letzte Aktivitäten",
	"Sign Out":                       "Abmelden",
	"Order History":                  "Auftragshistorie",
	"load from file":                 "von Datei laden",
	"loaded from file":               "von Datei geladen",
	"defaults":                       "Standardwerte",
	"Wallet Password":                "Wallet Passwort",
	"w_password_helper":              "Dies ist das Passwort das in der Walletsoftware konfiguriert wurde",
	"w_password_tooltip":             "Wenn für das Wallet kein Passwort erforderlich ist lass dieses Feld leer",
	"App Password":                   "App-Passwort",
	"Add":                            "Hinzufügen",
	"Unlock":                         "Entsperren",
	"Rescan":                         "Neu scannen",
	"Wallet":                         "Wallet",
	"app_password_reminder":          "Dein App-Passwort wird immer bei der Durchführung sensibler Walletoperationen benötigt.",
	"DEX Address":                    "DEX Adresse",
	"TLS Certificate":                "TLS Zertifikat",
	"remove":                         "Entfernen",
	"add a file":                     "Datei hinzufügen",
	"Submit":                         "Senden",
	"Confirm Registration":           "Registration bestätigen",                                                                    // stale but better than no translation
	"app_pw_reg":                     "Gib dein App-Passwort ein um die DEX-Registrierung zu bestätigen.",                          // stale...
	"reg_confirm_submit":             `Wenn du dieses Formular abschickst wird die Anmeldegebühr von deinem Geldbeutel abgebucht.`, // stale...
	"provided_markets":               "Dieser DEX bietet folgende Märkte:",
	"accepted_fee_assets":            "Dieser DEX akzeptiert die folgenden Gebühren:",
	"base_header":                    "Basis",
	"quote_header":                   "Angebot",
	"lot_size_headsup":               "Alle Trades werden in Vielfachen der Lotgröße getätigt.",
	"Password":                       "Passwort",
	"Register":                       "Registrieren",
	"Authorize Export":               "Export genehmigen",
	"export_app_pw_msg":              "Gib dein App-Passwort ein und genehmige den Kontoexport für",
	"Disable Account":                "Konto deaktivieren",
	"disable_app_pw_msg":             "Gib dein App-Passwort ein um dein Konto zu deaktivieren.",
	"disable_dex_server":             "Dieser DEX-Server kann jederzeit wieder aktiviert werden, indem Du ihn erneut hinzufügst.",
	"Authorize Import":               "Import genehmigen",
	"app_pw_import_msg":              "Gib dein App-Passwort ein um den Kontoimport zu bestätigen.",
	"Account File":                   "Konto Datei",
	"Change Application Password":    "App-Passwort ändern",
	"Current Password":               "Aktuelles Passwort",
	"New Password":                   "Neues Passwort",
	"Confirm New Password":           "Neues Passwort bestätigen",
	"cancel_no_pw":                   "Sende einen Stornierungsauftrag für die verbleibenden",
	"cancel_remain":                  "Der verbleibende Betrag kann sich ändern bevor der Stornierungsauftrag ausgeführt wird.",
	"Log In":                         "Anmelden",
	"epoch":                          "Epoche",
	"price":                          "Preis",
	"volume":                         "Volumen",
	"buys":                           "Käufe",
	"sells":                          "Verkäufe",
	"Buy Orders":                     "Kaufaufträge",
	"Quantity":                       "Menge",
	"Rate":                           "Kurs",
	"Limit Order":                    "Limit-Auftrag",
	"Market Order":                   "Markt-Auftrag",
	"reg_status_msg":                 `Um bei <span id="regStatusDex" class="text-break"></span> handeln zu können muss die Zahlung der Anmeldegebühr <span id="confReq"></span> bestätigt werden.`,
	"Buy":                            "Kaufen",
	"Sell":                           "Verkaufen",
	"lot_size":                       "Lotgröße",
	"Rate Step":                      "Kurs Stufe",
	"Max":                            "Max",
	"lot":                            "Lot",
	"min trade is about":             "min trade ist ungefähr",
	"immediate_explanation":          "Wird der Auftrag im nächsten Matchingzyklus nicht vollständig gedeckt wird die ungedeckte Menge weder nachgebucht noch erneut gematcht. Nur-Käufer-Auftrag.",
	"Immediate or cancel":            "Sofort oder Abbrechen",
	"Balances":                       "Guthaben",
	"outdated_tooltip":               "Der Guthabenstand ist möglicherweise veraltet. Stell eine Verbindung zum Wallet her um ihn zu aktualisieren.",
	"available":                      "verfügbar",
	"connect_refresh_tooltip":        "Zum Verbinden und Aktualisieren klicken",
	"add_a_base_wallet":              `Füge ein<br><span data-unit="base"></span><br>Wallet hinzu`,
	"add_a_quote_wallet":             `Füge ein<br><span data-unit="quote"></span><br>Wallet hinzu`,
	"locked":                         "gesperrt",
	"immature":                       "unfertig",
	"Sell Orders":                    "Verkaufsaufträge",
	"Your Orders":                    "Deine Aufträge",
	"Type":                           "Typ",
	"Side":                           "Seite",
	"Age":                            "Alter",
	"Filled":                         "Erfüllt",
	"Settled":                        "Abgewickelt",
	"Status":                         "Status",
	"view order history":             "Auftragshistorie anzeigen",
	"cancel_order":                   "auftrag abbrechen",
	"order details":                  "Auftragsdetails",
	"verify_order":                   `Verifiziere Auftrag <span id="vSideHeader"></span>`,
	"You are submitting an order to": "Du sendest ein Auftrag an",
	"at a rate of":                   "zu einer Rate von",
	"for a total of":                 "für eine Gesamtsumme von",
	"verify_market":                  "Es handelt sich hierbei um einen Marktauftrag der mit den besten verfügbaren Aufträgen im Buch gematcht wird. Auf der Grundlage des aktuellen Mid-Gap-Kurses auf dem Markt erhältst Du voraussichtlich etwa",
	"auth_order_app_pw":              "Autorisiere diesen Auftrat mit deinem App-Passwort.",
	"lots":                           "Lots",
	"order_disclaimer": `<span class="red">WICHTIG</span>: Der Handel braucht Zeit zur Abwicklung und du darfst währenddessen den DEX Klient   
		oder die<span data-unit="quote"></span> oder <span data-unit="base"></span> Blockchain und/oder die Walletsoftware nicht schließen bis
		die Abwicklung abgeschlossen ist. Die Abwicklung kann innerhalb weniger Minuten abgeschlossen sein oder bis zu einigen Stunden dauern.`,
	"Order":                     "Auftrag",
	"see all orders":            "Alle Aufträge anzeigen",
	"Exchange":                  "Exchange",
	"Market":                    "Markt",
	"Offering":                  "Angebot",
	"Asking":                    "Preisvorstellung",
	"Fees":                      "Gebühren",
	"order_fees_tooltip":        "On-Chain-Transaktionsgebühren die normalerweise vom Miner erhoben werden. Decred DEX verlangt keine Handelsgebühren.",
	"Matches":                   "Matches",
	"Match ID":                  "Match ID",
	"Time":                      "Zeit",
	"ago":                       "zuvor",
	"Cancellation":              "Stornierung",
	"Order Portion":             "Auftragsanteil",
	"you":                       "Du",
	"them":                      "Gegenseite",
	"Redemption":                "Rücknahme",
	"Refund":                    "Erstattung",
	"Funding Coins":             "Funding-Coins",
	"Exchanges":                 "Exchanges",
	"apply":                     "Anwenden",
	"Assets":                    "Assets",
	"Trade":                     "Handeln",
	"Set App Password":          "App-Passwort festlegen",
	"reg_set_app_pw_msg":        "Vergebe dein App-Passwort. Dieses Passwort schützt deine DEX-Kontoschlüssel und die angeschlossenen Wallets.",
	"Password Again":            "Passwort wiederholen",
	"Add a DEX":                 "Einen DEX hinzufügen",
	"Pick a server":             "Server auswählen",
	"reg_ssl_needed":            "Es sieht so aus als ob wir kein SSL-Zertifikat für diesen DEX haben. Füge das Zertifikat des Servers hinzu um fortzufahren.",
	"Dark Mode":                 "Dark Mode",
	"Show pop-up notifications": "Popup-Benachrichtigungen anzeigen",
	"Account ID":                "Account ID",
	"Export Account":            "Account exportieren",
	"simultaneous_servers_msg":  "Der Decred DEX Client unterstützt die gleichzeitige Nutzung einer beliebigen Anzahl von DEX-Servern.",
	"Change App Password":       "App-Passwort ändern",
	"Build ID":                  "Build ID",
	"Connect":                   "Verbinden",
	"Send":                      "Senden",
	"Deposit":                   "Einzahlen",
	"Lock":                      "Sperren",
	"New Deposit Address":       "Neue Einzahlungsadresse",
	"Address":                   "Adresse",
	"Amount":                    "Betrag",
	"Authorize the transfer with your app password.": "Autorisiere die Transaktion mit deinem App-Passwort.",
	"Reconfigure":                 "Rekonfigurieren",
	"pw_change_instructions":      "Das Ändern des Passworts ändert nicht das Passwort für deine Wallet-Software. Verwende dieses Formular um den DEX-Client zu aktualisieren nachdem du dein Passwort direkt in der Wallet-Software geändert hast.",
	"New Wallet Password":         "Neues Wallet-Passwort",
	"pw_change_warn":              "Hinweis: Der Wechsel zu einer anderen Geldbörse während eines aktiven Handels kann zum Verlust von Guthaben führen.",
	"Show more options":           "Weitere Optionen anzeigen",
	"seed_implore_msg":            "Du solltest deinen App-Seed sorgfältig notieren und eine Kopie davon speichern. Solltest du den Zugriff auf diesen Rechner oder die wichtigen Anwendungsdateien verlieren kann der Seed verwendet werden um deine DEX-Konten und nativen Wallets wiederherzustellen. Einige ältere Konten können nicht aus dem Seed wiederhergestellt werden. Unabhängig davon ob es sich um alte oder neue Konten handelt ist es eine empfehlenswerte Vorgehensweise die Kontoschlüssel getrennt vom Seed zu sichern.",
	"View Application Seed":       "App-Seed anzeigen",
	"Remember my password":        "Passwort merken",
	"pw_for_seed":                 "Gib dein App-Passwort ein um deinen Seed anzuzeigen. Stelle sicher dass niemand sonst deinen Bildschirm sehen kann.",
	"Asset":                       "Asset",
	"Balance":                     "Guthaben",
	"Actions":                     "Aktionen",
	"Restoration Seed":            "Wiederherstellungs-Seed",
	"Restore from seed":           "Aus Seed wiederherstellen",
	"Import Account":              "Account Importieren",
	"no_wallet":                   "kein Wallet",
	"create_a_x_wallet":           "Erstelle ein <span data-asset-name=1></span> Wallet",
	"dont_share":                  "Teile es nicht. Verliere es nicht.",
	"Show Me":                     "Anzeigen",
	"Wallet Settings":             "Wallet Einstellungen",
	"add_a_x_wallet":              `Füge ein <img data-tmpl="assetLogo" class="asset-logo mx-1"> <span data-tmpl="assetName"></span> Wallet hinzu`,
	"ready":                       "fertig",
	"off":                         "Ausgeschaltet",
	"Export Trades":               "Exportiere Trades",
	"change the wallet type":      "den Wallet-Typ ändern",
	"confirmations":               "Bestätigungen",
	"how_reg":                     "Wie soll die Anmeldegebühr bezahlt werden?",
	"All markets at":              "Alle Märkte bei",
	"pick a different asset":      "ein anderes Asset wählen",
	"Create":                      "Erstellen",
	"Register_loudly":             "Registrieren!",
	"1 Sync the Blockchain":       "1: Blockchain synchronisieren",
	"Progress":                    "Fortschritt",
	"remaining":                   "verbleibend",
	"2 Fund the Registration Fee": "2: Die Anmeldegebühr bezahlen",
	"One time anti-spam measure":  "Dies ist eine kleine, einmalige Anti-Spam-Maßnahme um störendes Verhalten wie z.B. das Zurückziehen von Swaps zu verhindern.",
	"Registration fee":            "Anmeldegebühr",
	"Your Deposit Address":        "Einzahlungsadresse für Ihr Wallet",
	// "Send enough for bonds":       `Achte darauf genug senden um auch die Netzwerkgebühren zu decken. Du kannst so viel einzahlen wie du willst da im nächsten Schritt nur die Registrierungsgebühr abgezogen wird. Du musst die Bestätigungen deiner Einzahlung abwarten damit du fortfahren kannst.`,
	"Send enough with estimate":   `Zahle mindestens <span data-tmpl="totalForBond"></span> <span class="unit">XYZ</span> ein um auch die Netzwerkgebühren der Transaktion zu decken. Du kannst so viel einzahlen wie du möchtest da nur der erforderliche Betrag im nächsten Schritt verrechnet wird. Du musst die Bestätigungen deiner Einzahlung abwarten damit du fortfahren kannst.`,
	"Send funds for token":        `Zahle mindestens <span data-tmpl="tokenFees"></span> <span class="unit">XYZ</span> und <span data-tmpl="parentFees"></span> <span data-tmpl="parentUnit">XYZ</span> ein um die Registrationsgebühr zu begleichen. Du kannst soviel in dein Wallet einzahlen wie du möchtest da nur der genannte Betrag für die Zahlung der Registrationsgebühr genutzt wird. Die Einzahlung muss erst bestätigt werden damit du fortfahren kannst.`,
	"add a different server":      "einen anderen Server hinzufügen",
	"Add a custom server":         "Benutzerdefinierten Server hinzufügen",
	"plus tx fees":                "+ tx fees",
	"Export Seed":                 "Seed exportieren",
	"Total":                       "Gesamt",
	"Trading":                     "Handeln",
	"Receiving Approximately":     "Empfange ungefähr",
	"Fee Projection":              "Gebührenprognose",
	"details":                     "Details",
	"to":                          "nach",
	"Options":                     "Optionen",
	"fee_projection_tooltip":      "Wenn sich die Netzwerkbedingungen nicht ändern bevor dein Auftrag erfüllt ist sollten die gesamten Gebühren innerhalb dieses Bereichs liegen.",
	"unlock_for_details":          "Entsperre dein Wallet um Auftragsdetails und zusätzliche Optionen zu erhalten.",
	"estimate_unavailable":        "Schätzungen und Optionen zum Auftrag nicht verfügbar",
	"Fee Details":                 "Gebühren Details",
	"estimate_market_conditions":  "Best- und Worst-Case-Schätzungen beruhen auf den aktuellen Netzwerkbedingungen und können sich bis zum Zeitpunkt der Auftragserfüllung ändern.",
	"Best Case Fees":              "Best-Case Gebühren",
	"best_case_conditions":        "Die Best-Case Gebühren fallen an wenn der gesamte Auftrag in einem einzigen Matching durchgeführt wird.",
	"Swap":                        "Swap",
	"Redeem":                      "Redeem",
	"Worst Case Fees":             "Worst-Case Gebühren",
	"worst_case_conditions":       "Der Worst-Case Fall kann eintreten wenn der Auftrag über viele Epochen hinweg in einzelnen Lots gematcht wird.",
	"Maximum Possible Swap Fees":  "Höchstmögliche Swap-Gebühren",
	"max_fee_conditions":          "Dies ist der höchste Betrag der für deinen Swap fällig werden könnte. Normalerweise betragen die Gebühren nur einen Bruchteil dieses Satzes. Der Höchstbetrag kann nicht mehr geändert werden sobald der Auftrag erteilt ist.",
	"wallet_logs":                 "Wallet Logs",
	"accelerate_order":            "Auftrag beschleunigen",
	"acceleration_text":           "Wenn deine Swap-Transaktionen stocken kannst du versuchen sie mit einer zusätzlichen Transaktion zu beschleunigen. Dies ist hilfreich wenn der Gebührensatz für eine bestehende unbestätigte Transaktion zu niedrig für das Mining im nächsten Block gewählt wurde, aber nicht wenn die Blöcke nur langsam gemined werden. Wenn du dieses Formular abschickst wird eine Transaktion erstellt die die Änderung aus deiner eigenen Swap-Initiierungstransaktion an dich selbst mit einer höheren Gebühr sendet. Der effektive Gebührensatz deiner Swap-Transaktionen wird zu dem Satz den du unten auswählst durchgeführt. Wähle einen Satz der ausreicht um in den nächsten Block aufgenommen zu werden. Ziehe einen Block-Explorer zu Rate um sicherzugehen.",
	"effective_swap_tx_rate":      "Effektive Swap tx Gebühr",
	"current_fee":                 "Derzeitig empfohlene Gebühr",
	"accelerate_success":          `Erfolgreich übermittelte Transaktion: <span id="accelerateTxID"></span>`,
	"accelerate":                  "Beschleunigen",
	"acceleration_transactions":   "Beschleunigungs-Transaktionen",
	"acceleration_cost_msg":       `Anhebung der effektiven Gebühr auf <span id="feeRateEstimate"></span> kostet <span id="feeEstimate"></span>`,
	"recent_acceleration_msg":     `Ihre letzte Beschleunigung ist erst <span id="recentAccelerationTime"></span> Minuten her! Sind Sie sicher das Sie beschleunigen wollen?`,
	"recent_swap_msg":             `Ihre älteste unbestätigte Swap-Transaktion wurde erst vor <span id="recentSwapTime"></span> Minuten eingereicht! Sind Sie sicher das Sie beschleunigen wollen?`,
	"early_acceleration_help_msg": `Es schadet deinem Auftrag nicht, aber du verschwendest möglicherweise Geld. Die Beschleunigung ist nur dann hilfreich wenn der Gebührensatz für eine bestehende unbestätigte Transaktion zu niedrig gewählt wurde um im nächsten Block gemined zu werden aber nicht wenn die Blöcke nur langsam gemined werden. Du kannst dies im Block-Explorer bestätigen indem Du dieses Popup schließt und auf deine vorherigen Transaktionen klickst.`,
	"recover":                     "Wiederherstellen",
	"recover_wallet":              "Wallet wiederherstellen",
	"recover_warning":             "Beim Wiederherstellen deines Wallet werden alle Wallet-Daten in einen Backup-Ordner verschoben. Du musst warten bis die Wallet wieder mit dem Netzwerk synchronisiert ist, was unter Umständen lange dauern kann bevor du die Wallet wieder benutzen kannst.",
	"wallet_actively_used":        "Wallet wird aktiv genutzt!",
	"confirm_force_message":       "Dieses Wallet verwaltet aktiv Aufträge. Nachdem du diese Aktion durchgeführt hast wird es sehr lange dauern bis deine Wallet wieder synchronisiert ist was dazu führen kann das die laufenden Aufträge fehlschlagen. Führe diese Aktion nur durch wenn es absolut notwendig ist!",
	"confirm":                     "Bestätigen",
	"cancel":                      "Abbrechen",
	"Update TLS Certificate":      "TLS Zertifikat aktualisieren",
	"registered dexes":            "Registrierte Dexes:",
	"successful_cert_update":      "Zertifikat wurde erfolgreich aktualisiert!",
	"update dex host":             "DEX Host aktualisieren",
	"copied":                      "Kopiert!",
	"export_wallet":               "Wallet exportieren",
	"pw_for_wallet_seed":          "Gib dein App-Passwort ein um das Wallet-Seed anzuzeigen. Stelle sicher dass niemand sonst deinen Bildschirm sehen kann. Wenn jemand Zugriff auf den Wallet-Seed erhält kann er dein gesamtes Guthaben stehlen.",
	"export_wallet_disclaimer":    `<span class="warning-text">Die Verwendung einer extern wiederhergestellten Wallet, während du aktive Trades im DEX laufen hast, kann zu fehlgeschlagenen Trades und VERLUST VON FINANZEN führen.</span> Es wird empfohlen das du deine Wallet nicht exportierst, es sei denn, du bist ein erfahrener Benutzer und weißt, was du tust.`,
	"export_wallet_msg":           "Nachfolgend findest du die erforderlichen Seeds um dein Wallet in einigen der beliebten externen Wallets wiederherzustellen. Führe KEINE Transaktionen mit deinen externen Wallet durch während du aktive Trades auf dem DEX laufen hast.",
	"clipboard_warning":           "Das Kopieren/Einfügen eines Wallet-Seeds stellt ein potenzielles Sicherheitsrisiko dar. Dies geschieht auf eigene Gefahr.",
	"fiat_exchange_rate_sources":  "Quellen für den Fiat-Wechselkurs",
	"Synchronizing":               "Synchronisieren",
	"wallet_wait_synced":          "Das Wallet wird nach der Synchronisierung erstellt",
	"Create a Wallet":             "Ein Wallet erstellen",
	"Receive":                     "Empfangen",
	"Wallet Type":                 "Wallet Typ",
	"Peer Count":                  "Anzahl Peers",
	"Sync Progress":               "Sync-Fortschritt",
	"Settings":                    "Einstellungen",
	"asset_name Markets":          "<span data-asset-name=1></span> Märkte",
	"Host":                        "Host",
	"No Recent Activity":          "Keine aktuelle Aktivität",
	"Recent asset_name Activity":  "<span data-asset-name=1></span>Aktuelle Aktivitäten",
	"other_actions":               "Andere Aktionen",
}
