# HorribleOrganizer

*So shoddy yet so organized.*

HorribleOrganizer renames and organizes videos downloaded from HorribleSubs so as to be consumed by Plex, Infuse, etc. Episode titles are retrieved from TheTVDB API and put into filenames for extra context. Especially useful for long-running shows where HorribleSubs and fans use total episode number yet media organization systems insist on dividing into seasons.

An example session:

```console
$ ./HorribleOrganizer /Volumes/Downloads/BitTorrent/\[HorribleSubs\]\ One\ Piece\ -\ *\ \[1080p\].mkv
'[HorribleSubs] One Piece - 867 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E088 - No.867 Lurking in the Darkness! Assassin Attacks Luffy!.mkv'
'[HorribleSubs] One Piece - 868 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E089 - No.868 A Man's Resolution! Katakuri's Life-Risking Great Match.mkv'
'[HorribleSubs] One Piece - 869 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E090 - No.869 Wake Up! To Cross over the Strongest Kenbunshoku.mkv'
'[HorribleSubs] One Piece - 870 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E091 - No.870 A God Speed Fist! New Gear 4 Activation!.mkv'
'[HorribleSubs] One Piece - 871 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E092 - No.871 Finally, It's Over! The Climax of the Intense Fight against Katakuri!.mkv'
'[HorribleSubs] One Piece - 872 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E093 - No.872 A Desperate Situation! The Iron-Tight Entrapment of Luffy!.mkv'
'[HorribleSubs] One Piece - 873 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E094 - No.873 Pulling Back from the Brink! The Formidable Reinforcements – Germa!.mkv'
'[HorribleSubs] One Piece - 874 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E095 - No.874 The Last Hope! The Sun Pirates Emerge!.mkv'
'[HorribleSubs] One Piece - 875 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E096 - No.875 A Captivating Flavor! Sanji's Cake of Happiness!.mkv'
'[HorribleSubs] One Piece - 876 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E097 - No.876 The Man of Humanity and Justice! Jinbe, a Desperate Massive Ocean Current.mkv'
'[HorribleSubs] One Piece - 877 [1080p].mkv'	=>	'One Piece/Season 19/One Piece - S19E098 - No.877 Time for Farewell! Pudding’s One Last Request!.mkv'
'[HorribleSubs] One Piece - 878 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E001 - No.878 The World in Shock! The Fifth Emperor of the Sea Arrives!.mkv'
'[HorribleSubs] One Piece - 879 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E002 - No.879 To the Reverie! Gathering of the Straw Hat Allies!.mkv'
'[HorribleSubs] One Piece - 880 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E003 - No.880 Sabo Goes into Action - All the Captains of the Revolutionary Army Appear!.mkv'
'[HorribleSubs] One Piece - 881 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E004 - No.881 The Next Move - New Obsessive Fleet Admiral Sakazuki.mkv'
'[HorribleSubs] One Piece - 882 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E005 - No.882 The Summit War - Pirate King's Inherited Will.mkv'
'[HorribleSubs] One Piece - 883 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E006 - No.883 One Step Ahead of the Dream - Shirahoshi's Path to the Sun!.mkv'
'[HorribleSubs] One Piece - 884 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E007 - No.884 I Miss Him! Vivi and Rebecca's Sentiments!.mkv'
'[HorribleSubs] One Piece - 885 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E008 - No.885 In the Dark Recesses of the Holyland! A Mysterious Giant Straw Hat!.mkv'
'[HorribleSubs] One Piece - 886 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E009 - No.886 The Holyland in Tumult! The Targeted Princess Shirahoshi!.mkv'
'[HorribleSubs] One Piece - 887 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E010 - No.887 An Explosive Situation! Two Emperors of the Sea Going After Luffy!.mkv'
'[HorribleSubs] One Piece - 888 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E011 - No.888 Sabo Enraged! The Tragedy of the Revolutionary Army Officer Kuma!.mkv'
'[HorribleSubs] One Piece - 889 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E012 - No.889 Finally, It Starts! The Conspiracy-Filled Reverie!.mkv'
'[HorribleSubs] One Piece - 890 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E013 - No.890 Marco! The Keeper of Whitebeard's Last Memento.mkv'
'[HorribleSubs] One Piece - 891 [1080p].mkv'	=>	'One Piece/Season 20/One Piece - S20E014 - No.891 Climbing Up a Waterfall! A Great Journey Through the Land of Wano's Sea Zone!.mkv'
'[HorribleSubs] One Piece - 892 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E001 - No.892 The Land of Wano! To the Samurai Country where Cherry Blossoms Flutter!.mkv'
'[HorribleSubs] One Piece - 893 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E002 - No.893 Otama Appears! Luffy vs. Kaido's Army!.mkv'
'[HorribleSubs] One Piece - 894 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E003 - No.894 He'll Come! The Legend of Ace in the Land of Wano!.mkv'
'[HorribleSubs] One Piece - 895 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E004 - No.895 Side Story! The World's Greatest Bounty Hunter, Cidre!.mkv'
'[HorribleSubs] One Piece - 896 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E005 - No.896 Side Story! Clash! Luffy vs. the King of Carbonation!.mkv'
'[HorribleSubs] One Piece - 897 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E006 - No.897 Save Otama! Straw Hat, Bounding through the Wasteland!.mkv'
'[HorribleSubs] One Piece - 898 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E007 - No.898 The Headliner! Hawkins the Magician Appears!.mkv'
'[HorribleSubs] One Piece - 899 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E008 - No.899 Defeat is Inevitable! The Strawman's Fierce Attack!.mkv'
'[HorribleSubs] One Piece - 900 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E009 - No.900 The Greatest Day of My Life! Otama and Her Sweet Red-bean Soup!.mkv'
'[HorribleSubs] One Piece - 901 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E010 - No.901 Charging into the Enemy's Territory! Bakura Town - Where Officials Thrive!.mkv'
'[HorribleSubs] One Piece - 902 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E011 - No.902 The Yokozuna Appears! The Invincible Urashima Goes After Okiku!.mkv'
'[HorribleSubs] One Piece - 903 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E012 - No.903 A Climactic Sumo Battle! Straw Hat vs. the Strongest Ever Yokozuna!.mkv'
'[HorribleSubs] One Piece - 904 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E013 - No.904 Luffy Rages! Rescue Otama from Danger!.mkv'
'[HorribleSubs] One Piece - 905 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E014 - No.905 Taking Back Otama! A Fierce Fight Against Holdem!.mkv'
'[HorribleSubs] One Piece - 906 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E015 - No.906 Duel! The Magician and the Surgeon of Death!.mkv'
'[HorribleSubs] One Piece - 907 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E016 - No.907 Romance Dawn.mkv'
'[HorribleSubs] One Piece - 908 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E017 - No.908 The Coming of the Treasure Ship! LuffytaroReturns the Favor!.mkv'
'[HorribleSubs] One Piece - 909 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E018 - No.909 Mysterious Grave Markers! A Reunion at the Ruins of Oden Castle!.mkv'
'[HorribleSubs] One Piece - 910 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E019 - No.910 A Legendary Samurai! The Man Who Roger Admired!.mkv'
'[HorribleSubs] One Piece - 911 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E020 - No.911 Bringing Down the Emperor of the Sea! A Secret Raid Operation Begins!.mkv'
'[HorribleSubs] One Piece - 912 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E021 - No.912 The Strongest Man in the World! Shutenmaru, the Thieves Brigade Chief!.mkv'
'[HorribleSubs] One Piece - 913 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E022 - No.913 Everyone is Annihilated! Kaido's Furious Blast Breath!.mkv'
'[HorribleSubs] One Piece - 914 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E023 - No.914 Finally Clashing! The Ferocious Luffy vs. Kaido!.mkv'
'[HorribleSubs] One Piece - 915 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E024 - No.915 Destructive! One Shot, One Kill： Thunder Bagua!.mkv'
'[HorribleSubs] One Piece - 916 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E025 - No.916 A Living Hell! Luffy, Humiliated in the Great Mine!.mkv'
'[HorribleSubs] One Piece - 917 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E026 - No.917 The Holyland in Tumult! Emperor of the Sea Blackbeard Cackles!.mkv'
'[HorribleSubs] One Piece - 918 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E027 - No.918 It's On! The Special Operation to Bring Down Kaido!.mkv'
'[HorribleSubs] One Piece - 919 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E028 - No.919 Rampage! The Prisoners： Luffy and Kid!.mkv'
'[HorribleSubs] One Piece - 920 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E029 - No.920 A Great Sensation! Sanji's Special Soba!.mkv'
'[HorribleSubs] One Piece - 921 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E030 - No.921 Luxurious and Gorgeous! Wano's Most Beautiful Woman： Komurasaki!.mkv'
'[HorribleSubs] One Piece - 922 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E031 - No.922 A Tale of Chivalry! Zoro and Tonoyasu's Little Trip!.mkv'
'[HorribleSubs] One Piece - 923 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E032 - No.923 A State of Emergency! Big Mom Closes In!.mkv'
'[HorribleSubs] One Piece - 924 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E033 - No.924 The Capital in an Uproar! Another Assassin Targets Sanji!.mkv'
'[HorribleSubs] One Piece - 925 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E034 - No.925 Dashing! The Righteous Soba Mask!.mkv'
'[HorribleSubs] One Piece - 926 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E035 - No.926 A Desperate Situation! Orochi's Menacing Oniwabanshu!.mkv'
'[HorribleSubs] One Piece - 927 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E036 - No.927 Pandemonium! The Monster Snake, Shogun Orochi!.mkv'
'[HorribleSubs] One Piece - 928 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E037 - No.928 The Flower Falls! The Final Moment of the Most Beautiful Woman in the Land of Wano!.mkv'
'[HorribleSubs] One Piece - 929 [1080p].mkv'	=>	'One Piece/Season 21/One Piece - S21E038 - No.929 The Bond Between Prisoners! Luffy and Old Man Hyo!.mkv'
```

## Installation and usage

Clone the repository, then

```
go build
cp config.example.toml config.toml
```

Edit `config.toml` to define shows to be organized. Then run `./HorribleOrganizer` to get started.

On first use you'll be prompted for a TheTVDB API key, registered at <https://thetvdb.com/dashboard/account/apikey>. It will be saved along with a periodically refreshed token in `auth.toml`.
