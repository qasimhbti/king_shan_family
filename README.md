

Coding challenge on GeekTrust, link: https://www.geektrust.in/coding-problem/backend/family


VERSION 1.0:
============

	LANGUAGE USED: GoLang (Version 1.11.4).

	TO RUN:
		1. Extract all the files into a new folder named goworkspace(bin, pkg, src).
		2. GOPATH should be set.
		2. Open a terminal and cd to this folder.
		3. Run the follwing commands:
			go build
			./family
		4. Follow the instructions provided by the program at each step.
		5. Problem 4. (Who's Your Daddy) is not implemented yet, will be available in version V 2.0.

		6. Family Member names and relation should be enter as follws:-
			KingShan, QueenAnga, Ish, Chit, Ambi, Vich, Lika, Satya, Vyan, Drita, Jaya, Vrita, Vila, Jnki, Chika, Kpila, Satvy, Asva, Savya, Krpi, Saayan, Mina, Jata, Driya, Mnu, Lavanya, Gru, Kriya, Misa  &&&
			Brother/Brothers, Sister/Sisters, Father, Mother, Children, Son/Sons, Daughter/Daughters, Cousins, GrandDaughter, GrandSon, GrandChildren, Brotherinlaw, Sisterinlaw, PaternalAunts/PaternalAunt, MaternalAunts/MaternalAunt, PaternalUncles/PaternalUncle, MaternalUncles/MaternalUncle.


	
	DATA STRUCTURE:
		1. Used a Binary Tree structure to represent King Shan Famiy.
		2. Planning to modify tree structure in future version (V 2.0).

	MODELLING:
		1. All the coding is done using Object Oriented Modelling using Struct and Method properties of GoLang.

	TESTS:
		1. Tests cases coding is in progress, will be available in version V 2.0.
                                                                                    
	MISC:
		1. Family Tree structure is explained in family_tree.json.
		2. Every family member details is as follows:-
			          
			a) id: Unique integer number starting from 2 ( 1 is preserved for parent of King Shan and Queen Anga)
			        
			b) parentID: which represent id of parent. Thus id = parentID of King Shan & Queen Anga = 1.
                   
			c) name: member name
                    
			d) gender: member gender
                     
			e) spouse: member spouse name
                      
			f) spouse Gender: spouse gender
               
			g) spouse parentID: currently all are 0 bacuase in the given family detail, parent of spouse are not provivded.

			
FEATURES PLANNED IN NEXT VERSION 2.0:
=====================================
		
		1. Problem 4. (Who's Your Daddy) is not implemented yet, will be available in version V 2.0.
		2. Tests cases coding is in progress, will be available in version V 2.0.
		3. Planning to modify tree structure in future version (V 2.0).

			NEW STRUCTURE 

											King-Shan & Queen Anga    ------------------------  ROOT NODE 
                                 				 |         |
                                			|  	             	|  
                            			|							|
                        			|									|	
             			Son & his Spouse Details            Daughter & her Spouse Details
                        			|											|
                    			|		|									|		|
                			|				|							|				|
            			|						|											|
        			|								|				|							|
Son & his Spouse Details  Daughter & her Spouse Details  Son & his Spouse Details  Daughter & herSpouse 
			|							|								|						|	
		|		|					|		|						|		|				|		|
	|				|			|				|				|				|		|				|

